package qiniu

import (
	"github.com/curder/go-qiniu-demo/pkg/log"
	qiniuService "github.com/curder/go-qiniu-demo/service/filesystem/qiniu"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
)

var privateBucketConfig = map[string]string{
	"name":      "test-web-private",
	"accessKey": "",
	"secretKey": "",
	"endPoint":  "http://qb4lomabw.bkt.clouddn.com",
}

// 存储
func Store(c *gin.Context) {
	var (
		// accountModel *model.AccountModel

		storage *qiniuService.Client
		// err     error
	)

	storage = qiniuService.New(&qiniuService.Config{
		AccessKey:     privateBucketConfig["accessKey"],
		SecretKey:     privateBucketConfig["secretKey"],
		Bucket:        privateBucketConfig["name"],
		Endpoint:      privateBucketConfig["endPoint"],
		UseHTTPS:      false,
		UseCdnDomains: false,
		PrivateURL:    true,
	})

	list(storage)

	// getUrl(c, storage)

	// delete(c, storage, "WX20200529-210325@2x.png")

	//put(c, storage)
}

// 文件列表
func list(storage *qiniuService.Client) {
	var (
		fileList   []*qiniuService.Object
		fileObject *qiniuService.Object
		err        error
	)
	if fileList, err = storage.List("/"); err != nil {
		panic(err)
		return
	}

	for _, fileObject = range fileList {
		file, _ := fileObject.StorageInterface.GetURL(fileObject.Path)
		log.Infof("%#v", file)
	}
}

// 根据文件名称获取地址
func getUrl(c *gin.Context, storage *qiniuService.Client) {
	var (
		url string
		err error
	)
	if url, err = storage.GetURL("qiniu_do_not_delete.gif"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"url": url,
	})
}

// 删除
func delete(c *gin.Context, storage *qiniuService.Client, path string) {
	if err := storage.Delete(path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err,
		})
	}

	c.JSON(http.StatusOK, nil)
}

//
func put(c *gin.Context, s *qiniuService.Client) {
	var (
		err        error
		pathPrefix string
		formData   *multipart.Form // *multipart.Form
		file       multipart.File
		// files map[string]*multipart.FileHeader
	)
	// 设置文件大小
	if err = c.Request.ParseMultipartForm(4 << 20); err != nil { // 4 MB
		c.JSON(http.StatusBadRequest, gin.H{"msg": "文件太大"})
		return
	}
	formData = c.Request.MultipartForm

	pathPrefix = c.Request.Form.Get("path_prefix")
	files := formData.File["file"]

	for _, v := range files {
		if file, err = v.Open(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "文件打开失败"})
			return
		}
		defer file.Close()

		if _, err := s.Put(pathPrefix+"5.png", file); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err})

			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"msg": "上传成功"})
}
