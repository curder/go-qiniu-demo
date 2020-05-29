package users

import (
	"github.com/curder/go-qiniu-demo/handler"
	"github.com/curder/go-qiniu-demo/model"
	"github.com/curder/go-qiniu-demo/pkg/errno"
	"github.com/curder/go-qiniu-demo/pkg/log"
	"github.com/curder/go-qiniu-demo/service/user"
	"github.com/gin-gonic/gin"
)

// 查询当前登录的用户信息
func Info(c *gin.Context) {
	var (
		userID    uint64
		userModel *model.UserModel
		err       error
	)
	log.Info("User Info function called.")

	userID = handler.GetUserID(c)

	// 通过用户ID查找用户信息
	if userModel, err = user.UserSvc.GetByID(userID); err != nil {
		handler.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	handler.SendResponse(c, nil, model.UserInfo{
		ID:        userModel.ID,
		Username:  userModel.Username,
		Email:     userModel.Email,
		CreatedAt: userModel.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: userModel.UpdatedAt.Format("2006-01-02 15:04:05"),
	})
}
