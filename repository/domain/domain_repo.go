package domain

import (
	"github.com/curder/go-qiniu-demo/model"
	"github.com/jinzhu/gorm"
)

// Repo 定义域名仓库接口
type Repo interface {
	Create(db *gorm.DB, domain model.DomainModel) (id uint64, err error)
	Delete(db *gorm.DB, id uint64) (rowsAffected int64, err error)
	Restore(db *gorm.DB, id uint64) (RowsAffected int64, err error)
	ForceDelete(db *gorm.DB, id uint64) (rowsAffected int64, err error)
	Update(db *gorm.DB, id uint64, domainMap map[string]interface{}) error
	GetDomains(db *gorm.DB) (domains []*model.DomainModel, err error)
	GetDomainByID(db *gorm.DB, id uint64) (domain *model.DomainModel, err error)
}

// domainRepo 域名仓库
type domainRepo struct{}

// 创建域名
func (repo *domainRepo) Create(db *gorm.DB, domain model.DomainModel) (id uint64, err error) {
	if err = db.Create(&domain).Error; err != nil {
		return 0, err
	}

	return domain.ID, nil
}

// 删除域名
func (repo *domainRepo) Delete(db *gorm.DB, id uint64) (rowsAffected int64, err error) {
	var (
		domain model.DomainModel
		result *gorm.DB
	)

	result = db.Where("id = ?", id).Delete(&domain)

	return result.RowsAffected, result.Error
}

// 恢复域名
func (repo *domainRepo) Restore(db *gorm.DB, id uint64) (RowsAffected int64, err error) {
	var (
		account model.AccountModel
		result  *gorm.DB
	)

	result = db.Unscoped().Where("id = ?", id).Find(&account).Update("deleted_at", gorm.Expr("NULL"))

	return result.RowsAffected, result.Error
}

// 强制删除域名
func (repo *domainRepo) ForceDelete(db *gorm.DB, id uint64) (rowsAffected int64, err error) {
	var (
		domain model.DomainModel
		result *gorm.DB
	)

	result = db.Unscoped().Where("id = ?", id).Delete(&domain)

	return result.RowsAffected, result.Error
}

// 更新域名
func (repo *domainRepo) Update(db *gorm.DB, id uint64, domainMap map[string]interface{}) (err error) {
	var (
		domain *model.DomainModel
	)
	// 检查账户是否存在
	if domain, err = repo.GetDomainByID(db, id); err != nil {
		return err
	}

	return db.Model(domain).Updates(domainMap).Error
}

// 获取域名列表
func (repo *domainRepo) GetDomains(db *gorm.DB) (domains []*model.DomainModel, err error) {
	var (
		result *gorm.DB
	)

	domains = make([]*model.DomainModel, 0)
	result = db.Find(&domains)

	return domains, result.Error
}

// 通过ID获取域名信息
func (repo *domainRepo) GetDomainByID(db *gorm.DB, id uint64) (*model.DomainModel, error) {
	var (
		domain model.DomainModel
		result *gorm.DB
	)

	result = db.Where("id = ?", id).First(&domain)

	return &domain, result.Error
}

// NewDomainRepo 实例化存储桶仓库
func NewDomainRepo() Repo {
	return &domainRepo{}
}
