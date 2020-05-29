package user

import (
	"github.com/curder/go-qiniu-demo/model"
	"github.com/curder/go-qiniu-demo/pkg/auth"
	"github.com/curder/go-qiniu-demo/repository/user"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// Service 用户服务接口定义
// 使用大写的service对外保留方法
type Service interface {
	Create(user model.UserModel) (id uint64, err error)
	GetByID(id uint64) (*model.UserModel, error)
	GetByName(name string) (*model.UserModel, error)
	VerifyLogin(userModel *model.UserModel, password string) (*model.UserModel, error)
	VerifyUsernameOrEmailExists(username, email string) bool
}

// UserSvc 直接初始化，可以避免在使用时再实例化
var UserSvc = NewUserService()

// 用小写的 service 实现接口中定义的方法
type userService struct {
	userRepo user.Repo
}

// NewUserService 实例化一个userService
// 通过 NewService 函数初始化 Service 接口
// 依赖接口，不要依赖实现，面向接口编程
func NewUserService() Service {
	return &userService{
		userRepo: user.NewUserRepo(),
	}
}

// GetByID 通过ID查找用户信息
func (srv *userService) GetByID(id uint64) (*model.UserModel, error) {
	userModel, err := srv.userRepo.GetByID(model.GetDB(), id)
	if err != nil {
		return userModel, errors.Wrapf(err, "get user info err from db by id: %d", id)
	}

	return userModel, nil
}

// GetByName 通过用户名或者邮箱获取用户信息
func (srv *userService) GetByName(name string) (*model.UserModel, error) {
	var (
		userModel *model.UserModel
		err       error
	)

	if userModel, err = srv.userRepo.GetByName(model.GetDB(), name); err != nil || gorm.IsRecordNotFoundError(err) {
		return userModel, errors.Wrapf(err, "get user info err from db by email: %s", name)
	}

	return userModel, nil
}

// VerifyLogin 校验用户信息
func (srv *userService) VerifyLogin(userModel *model.UserModel, password string) (*model.UserModel, error) {
	var (
		err error
	)

	// 判断密码是否正确
	if err = auth.Compare(userModel.Password, password); err != nil {
		return userModel, err
	}

	return userModel, nil
}

// VerifyUsernameOrEmailExists 校验用户名或者邮箱是否存在
func (srv *userService) VerifyUsernameOrEmailExists(username, email string) bool {
	var (
		userModel *model.UserModel
		err       error
	)

	if userModel, err = srv.userRepo.VerifyUsernameOrEmailExists(model.GetDB(), username, email); err != nil {
		return false
	}

	if userModel.ID != 0 {
		return true
	}

	return false
}

// 创建用户
func (srv *userService) Create(user model.UserModel) (id uint64, err error) {
	// 加密用户密码
	if user.Password, err = auth.Encrypt(user.Password); err != nil {
		return 0, err
	}

	if id, err = srv.userRepo.Create(model.GetDB(), user); err != nil {
		return id, err
	}

	return id, nil
}
