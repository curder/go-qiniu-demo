package domain

import (
	"github.com/curder/go-qiniu-demo/model"
	"github.com/curder/go-qiniu-demo/repository/domain"
)

// Service 域名服务接口定义
// 使用大写的service对外保留方法
type Service interface {
	Create(domain model.DomainModel) (id uint64, err error)
	Delete(id uint64) (rowsAffected int64, err error)
	Restore(id uint64) (rowsAffected int64, err error)
	ForceDelete(id uint64) (rowsAffected int64, err error)
	Update(id uint64, domainMap map[string]interface{}) error
	GetList() (domains []*model.DomainModel, err error)
	GetByID(id uint64) (domain *model.DomainModel, err error)
}

// DomainSvc 直接初始化，可以避免在使用时再实例化
var DomainSvc = NewDomainService()

// 用小写的 service 实现接口中定义的方法
type domainSvc struct {
	domainRepo domain.Repo
}

// 创建域名
func (svc *domainSvc) Create(domain model.DomainModel) (id uint64, err error) {
	if id, err = svc.domainRepo.Create(model.GetDB(), domain); err != nil {
		return id, err
	}

	return id, nil
}

// 删除域名
func (svc *domainSvc) Delete(id uint64) (rowsAffected int64, err error) {
	if rowsAffected, err = svc.domainRepo.Delete(model.GetDB(), id); err != nil {
		return
	}

	return
}

// 恢复域名
func (svc *domainSvc) Restore(id uint64) (rowsAffected int64, err error) {
	if rowsAffected, err = svc.domainRepo.Restore(model.GetDB(), id); err != nil {
		return
	}

	return
}

// 强制删除域名
func (svc *domainSvc) ForceDelete(id uint64) (rowsAffected int64, err error) {
	if rowsAffected, err = svc.domainRepo.ForceDelete(model.GetDB(), id); err != nil {
		return
	}

	return
}

// 更新域名
func (svc *domainSvc) Update(id uint64, domainMap map[string]interface{}) (err error) {
	if err = svc.domainRepo.Update(model.GetDB(), id, domainMap); err != nil {
		return
	}
	return
}

// 获取域名列表
func (svc *domainSvc) GetList() (domains []*model.DomainModel, err error) {
	if domains, err = svc.domainRepo.GetList(model.GetDB()); err != nil {
		return
	}
	return
}

// 通过ID获取域名详情
func (svc *domainSvc) GetByID(id uint64) (domain *model.DomainModel, err error) {
	if domain, err = svc.domainRepo.GetByID(model.GetDB(), id); err != nil {
		return
	}
	return
}

// NewDomainService 实例化一个domainService
// 通过 NewService 函数初始化 Service 接口
// 依赖接口，不要依赖实现，面向接口编程
func NewDomainService() Service {
	return &domainSvc{
		domainRepo: domain.NewDomainRepo(),
	}
}
