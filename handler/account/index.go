package account

import (
	"github.com/curder/go-qiniu-demo/handler"
	"github.com/curder/go-qiniu-demo/model"
	"github.com/curder/go-qiniu-demo/pkg/errno"
	"github.com/curder/go-qiniu-demo/pkg/log"
	"github.com/curder/go-qiniu-demo/service/account"
	"github.com/gin-gonic/gin"
)

// 账户列表
func Index(c *gin.Context) {
	log.Info("Get account list function called.")

	var (
		accounts []*model.AccountModel
		err      error
	)
	if accounts, err = account.AccountSvc.GetList(); err != nil {
		handler.SendResponse(c, errno.ErrAccountNotFound, nil)
		return
	}

	handler.SendResponse(c, nil, accounts)
}
