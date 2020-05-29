package account

import (
	"github.com/curder/go-qiniu-demo/model"
	"github.com/curder/go-qiniu-demo/repository/account"
)

// Service 用户服务接口定义
// 使用大写的service对外保留方法
type Service interface {
	Create(account model.AccountModel) (id uint64, err error)
	Delete(id uint64) (rowsAffected int64, err error)
	Restore(id uint64) (rowsAffected int64, err error)
	ForceDelete(id uint64) (rowsAffected int64, err error)
	Update(id uint64, accountMap map[string]interface{}) error
	GetList() ([]*model.AccountModel, error)
	GetByID(id uint64) (*model.AccountModel, error)
	GetByEmail(email string) (*model.AccountModel, error)
}

// AccountSvc 直接初始化，可以避免在使用时再实例化
var AccountSvc = NewAccountService()

// userRepo 用户仓库
type accountSvc struct {
	accountRepo account.Repo
}

// NewAccountService 实例化一个userService
// 通过 NewService 函数初始化 Service 接口
// 依赖接口，不要依赖实现，面向接口编程
func NewAccountService() Service {
	return &accountSvc{
		accountRepo: account.NewAccountRepo(),
	}
}

// 创建账户
func (srv *accountSvc) Create(account model.AccountModel) (id uint64, err error) {
	if id, err = srv.accountRepo.Create(model.GetDB(), account); err != nil {
		return id, err
	}

	return id, nil
}

// 删除账户
func (srv *accountSvc) Delete(id uint64) (rowsAffected int64, err error) {
	if rowsAffected, err = srv.accountRepo.Delete(model.GetDB(), id); err != nil {
		return
	}

	return
}

// 恢复账号
func (srv *accountSvc) Restore(id uint64) (rowsAffected int64, err error) {
	if rowsAffected, err = srv.accountRepo.Restore(model.GetDB(), id); err != nil {
		return
	}

	return
}

// 强制删除账户
func (srv *accountSvc) ForceDelete(id uint64) (rowsAffected int64, err error) {
	if rowsAffected, err = srv.accountRepo.ForceDelete(model.GetDB(), id); err != nil {
		return
	}

	return
}

// 更新账户
func (srv *accountSvc) Update(id uint64, accountMap map[string]interface{}) (err error) {
	if err = srv.accountRepo.Update(model.GetDB(), id, accountMap); err != nil {
		return
	}
	return
}

// 获取账户列表
func (srv *accountSvc) GetList() (modelAccounts []*model.AccountModel, err error) {
	if modelAccounts, err = srv.accountRepo.GetList(model.GetDB()); err != nil {
		return
	}
	return
}

// 通过ID获取账户信息
func (srv *accountSvc) GetByID(id uint64) (accountModel *model.AccountModel, err error) {
	if accountModel, err = srv.accountRepo.GetByID(model.GetDB(), id); err != nil {
		return
	}
	return
}

// 通过邮箱获取账户信息
func (srv *accountSvc) GetByEmail(email string) (accountModel *model.AccountModel, err error) {
	if accountModel, err = srv.accountRepo.GetByEmail(model.GetDB(), email); err != nil {
		return
	}
	return
}
