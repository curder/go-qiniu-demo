package account

import (
	"github.com/curder/go-qiniu-demo/handler"
	"github.com/curder/go-qiniu-demo/model"
	"github.com/curder/go-qiniu-demo/pkg/errno"
	"github.com/curder/go-qiniu-demo/pkg/log"
	"github.com/curder/go-qiniu-demo/service/account"
	"github.com/gin-gonic/gin"
	"time"
)

// 新增账户
func Store(c *gin.Context) {
	var (
		a   model.AccountModel
		id  uint64
		err error
	)

	log.Info("account store function called.")

	var req StoreAccountRequest
	if err = c.ShouldBind(&req); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	log.Infof("req %#v", req)

	a = model.AccountModel{
		Email:       req.Email,
		AccessKey:   req.AccessKey,
		SecretKey:   req.SecretKey,
		Description: req.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if id, err = account.AccountSvc.CreateAccount(a); err != nil {
		log.Warnf("[account] store account err, %v", err)
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}

	handler.SendResponse(c, nil, id)
}
