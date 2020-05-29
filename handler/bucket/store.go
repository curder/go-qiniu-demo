package bucket

import (
	"github.com/curder/go-qiniu-demo/handler"
	"github.com/curder/go-qiniu-demo/model"
	"github.com/curder/go-qiniu-demo/pkg/errno"
	"github.com/curder/go-qiniu-demo/pkg/log"
	"github.com/curder/go-qiniu-demo/service/account"
	"github.com/curder/go-qiniu-demo/service/bucket"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 保存存储桶
func Store(c *gin.Context) {
	var (
		b         model.BucketModel
		accountID int
		id        uint64
		err       error
	)

	log.Info("account store function called.")

	accountID, _ = strconv.Atoi(c.PostForm("account_id"))

	var req StoreBucketRequest
	if err = c.ShouldBind(&req); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	log.Infof("req %#v", req)

	// 判断账户是否存在
	if _, err = account.AccountSvc.GetByID(uint64(accountID)); err != nil {
		handler.SendResponse(c, errno.ErrAccountNotFound, nil)
		return
	}

	b = model.BucketModel{
		AccountID:   req.AccountID,
		Name:        req.Name,
		Description: req.Description,
	}

	if id, err = bucket.BucketSvc.Create(b); err != nil {
		log.Warnf("[bucket] store account err, %v", err)
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}

	handler.SendResponse(c, nil, id)
}
