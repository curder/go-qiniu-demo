package user

import (
	"github.com/curder/go-qiniu-demo/model"
	"github.com/jinzhu/gorm"
)

// Repo 定义用户仓库接口
type Repo interface {
	GetUserByID(db *gorm.DB, id uint64) (*model.UserModel, error)
	GetUserByName(db *gorm.DB, name string) (*model.UserModel, error)
	VerifyUsernameOrEmailExists(db *gorm.DB, username, email string) (*model.UserModel, error)
	Create(db *gorm.DB, user model.UserModel) (id uint64, err error)
}

// userRepo 用户仓库
type userRepo struct{}

// NewUserRepo 实例化用户仓库
func NewUserRepo() Repo {
	return &userRepo{}
}

// 通过用户ID获取用户信息
func (repo *userRepo) GetUserByID(db *gorm.DB, id uint64) (*model.UserModel, error) {
	user := &model.UserModel{}
	result := db.Where("id = ?", id).First(user)

	return user, result.Error
}

// GetUserByName 根据邮箱或用户名获取用户信息
func (repo *userRepo) GetUserByName(db *gorm.DB, name string) (*model.UserModel, error) {
	var (
		user   model.UserModel
		result *gorm.DB
	)

	result = db.Where("email = ?", name).Or("username = ?", name).First(&user)

	return &user, result.Error
}

// VerifyUsernameOrEmailExists 根据用户名和邮箱检查用户是否存在
func (repo *userRepo) VerifyUsernameOrEmailExists(db *gorm.DB, username, email string) (*model.UserModel, error) {
	var (
		user   model.UserModel
		result *gorm.DB
	)

	result = db.Where("email = ?", email).Or("username = ?", username).First(&user)

	return &user, result.Error
}

func (repo *userRepo) Create(db *gorm.DB, user model.UserModel) (id uint64, err error) {
	err = db.Create(&user).Error
	if err != nil {
		return 0, err
	}

	return user.ID, nil
}
