package users

import (
	"github.com/curder/go-qiniu-demo/handler"
	"github.com/curder/go-qiniu-demo/model"
	"github.com/curder/go-qiniu-demo/pkg/errno"
	"github.com/curder/go-qiniu-demo/pkg/log"
	"github.com/curder/go-qiniu-demo/pkg/token"
	"github.com/curder/go-qiniu-demo/service/user"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var (
		u   model.UserModel
		t   string
		err error
	)

	log.Info("User Register function called.")

	// Binding the data with the u struct.
	var req UserRegisterCredentials
	if err = c.ShouldBind(&req); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	log.Infof("req %#v", req)

	// 判断用户是否存在
	if user.UserSvc.VerifyUsernameOrEmailExists(req.Name, req.Email) {
		handler.SendResponse(c, errno.ErrUserExists, nil)
		return
	}

	// 创建用户
	u = model.UserModel{
		Username: req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	if u.ID, err = user.UserSvc.CreateUser(u); err != nil {
		log.Warnf("[register] create u err, %v", err)
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}

	// 签发签名 Sign the json web token.
	t, err = token.Sign(c, token.Context{UserID: u.ID, Username: u.Username, Email: u.Email}, "")
	if err != nil {
		log.Warnf("[register] gen token sign err:, %v", err)
		handler.SendResponse(c, errno.ErrToken, nil)
		return
	}

	handler.SendResponse(c, nil, model.Token{
		Token: t,
	})
}
