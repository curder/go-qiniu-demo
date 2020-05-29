package account

import (
	"github.com/curder/go-qiniu-demo/handler"
	"github.com/curder/go-qiniu-demo/pkg/errno"
	"github.com/curder/go-qiniu-demo/pkg/log"
	"github.com/curder/go-qiniu-demo/service/account"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 强制删除
func ForceDelete(c *gin.Context) {
	var (
		accountID    int
		rowsAffected int64
		err          error
	)

	log.Info("account force delete function called.")

	accountID, err = strconv.Atoi(c.Param("id"))

	if rowsAffected, err = account.AccountSvc.ForceDeleteAccount(uint64(accountID)); err != nil && rowsAffected != 0 {
		log.Warnf("[account] delete account err, %v", err)
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}

	handler.SendResponse(c, nil, accountID)
}
