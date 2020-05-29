package domain

import (
	"github.com/curder/go-qiniu-demo/handler"
	"github.com/curder/go-qiniu-demo/model"
	"github.com/curder/go-qiniu-demo/pkg/errno"
	"github.com/curder/go-qiniu-demo/pkg/log"
	"github.com/curder/go-qiniu-demo/service/bucket"
	"github.com/curder/go-qiniu-demo/service/domain"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 保存
func Store(c *gin.Context) {
	var (
		req      StoreDomainRequest
		d        model.DomainModel
		bucketID int
		id       uint64
		err      error
	)

	log.Info("domain store function called.")

	bucketID, _ = strconv.Atoi(c.PostForm("bucket_id"))

	if err = c.ShouldBind(&req); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	log.Infof("req %#v", req)

	// 判断存储桶是否存在
	if _, err = bucket.BucketSvc.GetBucketByID(uint64(bucketID)); err != nil {
		handler.SendResponse(c, errno.ErrBucketNotFound, nil)
		return
	}

	d = model.DomainModel{
		BucketID:    req.BucketID,
		Protocol:    req.Protocol,
		Hostname:    req.Hostname,
		Description: req.Description,
	}

	if id, err = domain.DomainSvc.Create(d); err != nil {
		log.Warnf("[domain] store domain err, %v", err)
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}

	handler.SendResponse(c, nil, id)
}
