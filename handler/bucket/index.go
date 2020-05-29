package bucket

import (
	"github.com/curder/go-qiniu-demo/handler"
	"github.com/curder/go-qiniu-demo/model"
	"github.com/curder/go-qiniu-demo/pkg/errno"
	"github.com/curder/go-qiniu-demo/pkg/log"
	"github.com/curder/go-qiniu-demo/service/bucket"
	"github.com/gin-gonic/gin"
)

// 存储桶列表
func Index(c *gin.Context) {
	log.Info("Get bucket list function called.")

	var (
		accounts []*model.BucketModel
		err      error
	)
	if accounts, err = bucket.BucketSvc.GetList(); err != nil {
		handler.SendResponse(c, errno.ErrAccountNotFound, nil)
		return
	}

	handler.SendResponse(c, nil, accounts)
}
