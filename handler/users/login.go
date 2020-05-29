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

// Login 用户登录接口
// @Summary 用户登录接口
// @Description 手机+密码登录
// @Tags 用户
// @Produce  json
// @Param req body PhoneLoginCredentials true "phone"
// @Success 200 {string} json "{"code":0,"message":"OK","data":{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6Ik"}}"
// @Router /auth/login [post]
func Login(c *gin.Context) {
	var (
		u   *model.UserModel
		t   string
		err error
	)
	log.Info("User Login function called.")
	// Binding the data with the u struct.
	var req UserLoginCredentials
	if err = c.ShouldBind(&req); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	log.Infof("req %#v", req)
	// check param
	if req.Name == "" || req.Password == "" {
		handler.SendResponse(c, errno.ErrParam, nil)
		return
	}

	// 如果用户已经存在，则通过用户名或者邮箱获取用户信息
	if u, err = user.UserSvc.GetByName(req.Name); err != nil {
		log.Warnf("[login] get u info err, %v", err)
		handler.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	// 验证用户密码
	if u, err = user.UserSvc.VerifyLogin(u, req.Password); err != nil {
		handler.SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}

	// 签发签名 Sign the json web token.
	t, err = token.Sign(c, token.Context{UserID: u.ID, Username: u.Username, Email: u.Email}, "")
	if err != nil {
		log.Warnf("[login] gen token sign err:, %v", err)
		handler.SendResponse(c, errno.ErrToken, nil)
		return
	}

	handler.SendResponse(c, nil, model.Token{
		Token: t,
	})
}
