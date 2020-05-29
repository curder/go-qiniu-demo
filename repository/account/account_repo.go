package account

import (
	"github.com/curder/go-qiniu-demo/model"
	"github.com/jinzhu/gorm"
)

// Repo 定义账户仓库接口
type Repo interface {
	Create(db *gorm.DB, user model.AccountModel) (id uint64, err error)
	Delete(db *gorm.DB, id uint64) (rowsAffected int64, err error)
	Restore(db *gorm.DB, id uint64) (RowsAffected int64, err error)
	ForceDelete(db *gorm.DB, id uint64) (rowsAffected int64, err error)
	Update(db *gorm.DB, id uint64, accountMap map[string]interface{}) error
	GetAccounts(db *gorm.DB) ([]*model.AccountModel, error)
	GetAccountByID(db *gorm.DB, id uint64) (*model.AccountModel, error)
	GetAccountByEmail(db *gorm.DB, email string) (*model.AccountModel, error)
}

// userRepo 用户仓库
type accountRepo struct{}

// NewAccountRepo 实例化用户仓库
func NewAccountRepo() Repo {
	return &accountRepo{}
}

// 创建账户接口
func (repo *accountRepo) Create(db *gorm.DB, account model.AccountModel) (id uint64, err error) {
	if err = db.Create(&account).Error; err != nil {
		return 0, err
	}

	return account.ID, nil
}

// 删除账户信息
func (repo *accountRepo) Delete(db *gorm.DB, id uint64) (rowsAffected int64, err error) {
	var (
		account model.AccountModel
		result  *gorm.DB
	)

	result = db.Where("id = ?", id).Delete(&account)

	return result.RowsAffected, result.Error
}

// 恢复删除的数据
func (repo *accountRepo) Restore(db *gorm.DB, id uint64) (RowsAffected int64, err error) {
	var (
		account model.AccountModel
		result  *gorm.DB
	)

	result = db.Unscoped().Where("id = ?", id).Find(&account).Update("deleted_at", gorm.Expr("NULL"))

	return result.RowsAffected, result.Error
}

// 强制删除账户信息
func (repo *accountRepo) ForceDelete(db *gorm.DB, id uint64) (rowsAffected int64, err error) {
	var (
		account model.AccountModel
		result  *gorm.DB
	)

	result = db.Unscoped().Where("id = ?", id).Delete(&account)

	return result.RowsAffected, result.Error
}

// 更新账户接口
func (repo *accountRepo) Update(db *gorm.DB, id uint64, accountMap map[string]interface{}) (err error) {
	var (
		account *model.AccountModel
	)
	// 检查账户是否存在
	if account, err = repo.GetAccountByID(db, id); err != nil {
		return err
	}

	return db.Model(account).Updates(accountMap).Error
}

// 获取账户列表
func (repo *accountRepo) GetAccounts(db *gorm.DB) (accounts []*model.AccountModel, error error) {
	var (
		result *gorm.DB
	)

	accounts = make([]*model.AccountModel, 0)
	result = db.Find(&accounts)

	return accounts, result.Error
}

// 通过ID获取账户信息
func (repo *accountRepo) GetAccountByID(db *gorm.DB, id uint64) (*model.AccountModel, error) {
	var (
		account model.AccountModel
		result  *gorm.DB
	)

	result = db.Where("id = ?", id).First(&account)

	return &account, result.Error
}

// 通过邮箱获取账户信息
func (repo *accountRepo) GetAccountByEmail(db *gorm.DB, email string) (*model.AccountModel, error) {
	var (
		account model.AccountModel
		result  *gorm.DB
	)

	result = db.Where("email = ?", email).First(&account)

	return &account, result.Error
}
