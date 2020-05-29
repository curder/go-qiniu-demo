package account

import (
	"github.com/curder/go-qiniu-demo/handler"
	"github.com/curder/go-qiniu-demo/pkg/errno"
	"github.com/curder/go-qiniu-demo/pkg/log"
	"github.com/curder/go-qiniu-demo/service/account"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 删除账户
func Delete(c *gin.Context) {
	var (
		accountID    int
		rowsAffected int64
		err          error
	)

	log.Info("account delete function called.")

	accountID, _ = strconv.Atoi(c.Param("id"))

	// 检查是否存在
	if _, err = account.AccountSvc.GetByID(uint64(accountID)); err != nil {
		log.Warnf("[account] delete account err, %v", err)
		handler.SendResponse(c, errno.ErrAccountNotFound, nil)
		return
	}

	if rowsAffected, err = account.AccountSvc.Delete(uint64(accountID)); err != nil && rowsAffected != 0 {
		log.Warnf("[account] delete account err, %v", err)
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}

	handler.SendResponse(c, nil, accountID)
}
