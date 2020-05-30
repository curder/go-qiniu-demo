package qiniu

import (
	"bytes"
	"context"
	"fmt"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// StorageInterface define common API to operate storage
type StorageInterface interface {
	Get(path string) (*os.File, error)
	GetStream(path string) (io.ReadCloser, error)
	Put(path string, reader io.Reader) (*Object, error)
	Delete(path string) error
	List(path string) ([]*Object, error)
	GetURL(path string) (string, error)
	GetEndpoint() string
}

// Object content object
type Object struct {
	Path             string
	Name             string
	LastModified     *time.Time
	StorageInterface StorageInterface
}

// Client Qiniu storage
type Client struct {
	Config        *Config
	mac           *qbox.Mac
	storageCfg    storage.Config
	bucketManager *storage.BucketManager
	putPolicy     *storage.PutPolicy
}

// Config Qiniu client config
type Config struct {
	AccessKey string
	SecretKey string
	// Region        string
	Bucket        string
	Endpoint      string
	UseHTTPS      bool
	UseCdnDomains bool
	PrivateURL    bool
}

// 自定义上传返回值结构体
type PutRet struct {
	Key    string
	Hash   string
	Fsize  int
	Bucket string
	Name   string
}

var urlRegexp = regexp.MustCompile(`(https?:)?//((\w+).)+(\w+)/`)

func New(config *Config) (client *Client) {
	var (
		z   *storage.Zone
		err error
	)
	client = &Client{Config: config, storageCfg: storage.Config{}}

	client.mac = qbox.NewMac(config.AccessKey, config.SecretKey)

	if z, err = getZone(config.AccessKey, config.Bucket); err != nil {
		panic(fmt.Sprintf("Zone is invalid, only support huadong, huabei, huanan, beimei, xinjiapo. err: %s", err))
	} else {
		client.storageCfg.Zone = z
	}

	//if len(config.Endpoint) == 0 {
	//	panic("endpoint must be provided.")
	//}
	if region, err := storage.GetRegion(config.AccessKey, config.Bucket); err != nil {
		client.storageCfg.Region = region
	}

	client.storageCfg.UseHTTPS = config.UseHTTPS
	client.storageCfg.UseCdnDomains = config.UseCdnDomains
	client.bucketManager = storage.NewBucketManager(client.mac, &client.storageCfg)

	return client
}

// 获取地区
func getZone(accessKey, bucket string) (zone *storage.Zone, err error) {
	if zone, err = storage.GetZone(accessKey, bucket); err != nil { // 空间对应的机房
		return
	}
	return
}

func (c *Client) SetPutPolicy(putPolicy *storage.PutPolicy) {
	c.putPolicy = putPolicy
}

// Get receive file with given path
func (c *Client) Get(path string) (file *os.File, err error) {
	var (
		readCloser io.ReadCloser
	)

	readCloser, err = c.GetStream(path)

	if file, err = ioutil.TempFile("/tmp", "qiniu"); err == nil {
		defer readCloser.Close()
		_, err = io.Copy(file, readCloser)
		file.Seek(0, 0)
	}

	return file, err
}

// GetStream get file as stream
func (c *Client) GetStream(path string) (readCloser io.ReadCloser, err error) {
	var (
		purl string
		res  *http.Response
	)
	if purl, err = c.GetURL(path); err != nil {
		return nil, err
	}

	res, err = http.Get(purl)
	if err == nil && res.StatusCode != http.StatusOK {
		err = fmt.Errorf("file %s not found", path)
	}

	readCloser = res.Body

	return
}

// Put store a reader into given path
func (c *Client) Put(urlPath string, reader io.Reader) (r *Object, err error) {
	var (
		seeker       io.Seeker
		ok           bool
		buffer       []byte
		fileType     string
		putPolicy    storage.PutPolicy
		upToken      string
		formUploader *storage.FormUploader
		ret          PutRet
		dataLen      int64
		putExtra     storage.PutExtra
		now          time.Time
	)

	if seeker, ok = reader.(io.ReadSeeker); ok {
		seeker.Seek(0, 0)
	}

	urlPath = storageKey(urlPath)
	if buffer, err = ioutil.ReadAll(reader); err != nil {
		return
	}

	fileType = mime.TypeByExtension(path.Ext(urlPath))
	if fileType == "" {
		fileType = http.DetectContentType(buffer)
	}

	putPolicy = storage.PutPolicy{
		Scope:      fmt.Sprintf("%s:%s", c.Config.Bucket, urlPath),
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`, // 使用 returnBody 自定义回复格式
	}

	if c.putPolicy != nil {
		putPolicy = *c.putPolicy
	}

	upToken = putPolicy.UploadToken(c.mac)

	formUploader = storage.NewFormUploader(&c.storageCfg)
	ret = PutRet{}
	dataLen = int64(len(buffer))

	putExtra = storage.PutExtra{
		Params: map[string]string{},
	}
	if err = formUploader.Put(context.Background(), &ret, upToken, urlPath, bytes.NewReader(buffer), dataLen, &putExtra); err != nil {
		return
	}

	now = time.Now()
	return &Object{
		Path:             ret.Key,
		Name:             filepath.Base(urlPath),
		LastModified:     &now,
		StorageInterface: c,
	}, err
}

// Delete delete file
func (c *Client) Delete(path string) error {
	return c.bucketManager.Delete(c.Config.Bucket, storageKey(path))
}

// Move move or rename file
func (c *Client) Move(path, destBucket, destKey string, force bool) (err error) {
	return c.bucketManager.Move(c.Config.Bucket, storageKey(path), destBucket, destKey, force)
}

// Fetch
func (c *Client) FetchRemoteResource(resURL, key string) (fetchRet storage.FetchRet, err error) {
	fetchRet, err = c.bucketManager.Fetch(resURL, c.Config.Bucket, key)
	return
}

// List list all objects under current path
func (c *Client) List(path string) (objects []*Object, err error) {
	var (
		prefix    = storageKey(path)
		listItems []storage.ListItem
		content   storage.ListItem
		t         time.Time
	)

	listItems, _, _, _, err = c.bucketManager.ListFiles(
		c.Config.Bucket,
		prefix,
		"",
		"",
		100,
	)

	if err != nil {
		return
	}

	for _, content = range listItems {
		t = time.Unix(content.PutTime, 0)
		objects = append(objects, &Object{
			Path:             "/" + storageKey(content.Key),
			Name:             filepath.Base(content.Key),
			LastModified:     &t,
			StorageInterface: c,
		})
	}

	return
}

// GetEndpoint get endpoint, FileSystem's endpoint is /
func (c *Client) GetEndpoint() string {
	return c.Config.Endpoint
}

func storageKey(urlPath string) string {
	var (
		u   *url.URL
		err error
	)

	if urlRegexp.MatchString(urlPath) {
		if u, err = url.Parse(urlPath); err == nil {
			urlPath = u.Path
		}
	}
	return strings.TrimPrefix(urlPath, "/")
}

// GetURL get public accessible URL
func (c *Client) GetURL(path string) (url string, err error) {
	var (
		key      string
		deadline int64
	)

	if len(path) == 0 {
		return
	}
	key = storageKey(path)

	if c.Config.PrivateURL {
		deadline = time.Now().Add(time.Second * 3600).Unix()
		url = storage.MakePrivateURL(c.mac, c.GetEndpoint(), key, deadline)
		return
	}

	url = storage.MakePublicURL(c.GetEndpoint(), key)

	return
}
