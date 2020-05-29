package account

import (
	"github.com/curder/go-qiniu-demo/handler"
	"github.com/curder/go-qiniu-demo/pkg/errno"
	"github.com/curder/go-qiniu-demo/pkg/log"
	"github.com/curder/go-qiniu-demo/service/account"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 恢复账号
func Restore(c *gin.Context) {
	var (
		accountID    int
		rowsAffected int64
		err          error
	)

	log.Info("account restore function called.")

	accountID, err = strconv.Atoi(c.Param("id"))

	if rowsAffected, err = account.AccountSvc.RestoreAccount(uint64(accountID)); err != nil && rowsAffected != 0 {
		log.Warnf("[account] delete account err, %v", err)
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}

	handler.SendResponse(c, nil, accountID)
}
