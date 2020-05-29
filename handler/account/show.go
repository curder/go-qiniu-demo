package account

import (
	"github.com/curder/go-qiniu-demo/handler"
	"github.com/curder/go-qiniu-demo/model"
	"github.com/curder/go-qiniu-demo/pkg/errno"
	"github.com/curder/go-qiniu-demo/pkg/log"
	"github.com/curder/go-qiniu-demo/service/account"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 账户详情
func Show(c *gin.Context) {
	var (
		accountID    int
		accountModel *model.AccountModel
		response     ShowAccountResponse
		err          error
	)

	log.Info("account show function called.")

	accountID, _ = strconv.Atoi(c.Param("id"))

	if accountModel, err = account.AccountSvc.GetAccountByID(uint64(accountID)); err != nil {
		log.Warnf("[account] show account err, %v", err)
		handler.SendResponse(c, errno.ErrAccountNotFound, nil)
		return
	}

	response = ShowAccountResponse{
		ID:          accountModel.ID,
		AccountKey:  accountModel.AccessKey,
		SecretKey:   accountModel.SecretKey,
		Description: accountModel.Description,
		CreatedAt:   accountModel.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   accountModel.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	handler.SendResponse(c, nil, response)
}
