package account

import (
	"github.com/curder/go-qiniu-demo/handler"
	"github.com/curder/go-qiniu-demo/pkg/errno"
	"github.com/curder/go-qiniu-demo/pkg/log"
	"github.com/curder/go-qiniu-demo/service/account"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 更新账户
func Update(c *gin.Context) {
	var (
		accountID  int
		accountMap map[string]interface{}
		err        error
	)

	log.Info("account update function called.")

	accountID, _ = strconv.Atoi(c.Param("id"))

	var req UpdateAccountRequest
	if err = c.ShouldBind(&req); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	log.Infof("req %#v", req)

	accountMap = make(map[string]interface{})
	accountMap["email"] = req.Email
	accountMap["access_key"] = req.AccessKey
	accountMap["secret_key"] = req.SecretKey
	accountMap["description"] = req.Description

	if err = account.AccountSvc.Update(uint64(accountID), accountMap); err != nil {
		log.Warnf("[account] store account err, %v", err)
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}

	handler.SendResponse(c, nil, accountID)
}
