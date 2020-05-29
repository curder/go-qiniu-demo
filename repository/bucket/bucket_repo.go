package bucket

import (
	"github.com/curder/go-qiniu-demo/model"
	"github.com/jinzhu/gorm"
)

// Repo 定义账户仓库接口
type Repo interface {
	Create(db *gorm.DB, bucket model.BucketModel) (id uint64, err error)
	Delete(db *gorm.DB, id uint64) (rowsAffected int64, err error)
	Restore(db *gorm.DB, id uint64) (RowsAffected int64, err error)
	ForceDelete(db *gorm.DB, id uint64) (rowsAffected int64, err error)
	Update(db *gorm.DB, id uint64, bucketMap map[string]interface{}) error
	GetBuckets(db *gorm.DB) (buckets []*model.BucketModel, err error)
	GetBucketByID(db *gorm.DB, id uint64) (bucket *model.BucketModel, err error)
	GetBucketByName(db *gorm.DB, name string) (bucket *model.BucketModel, err error)
}

// bucketRepo 用户仓库
type bucketRepo struct{}

// 创建存储桶接口
func (b bucketRepo) Create(db *gorm.DB, bucket model.BucketModel) (id uint64, err error) {
	if err = db.Create(&bucket).Error; err != nil {
		return 0, err
	}

	return bucket.ID, nil
}

// 删除存储桶
func (b bucketRepo) Delete(db *gorm.DB, id uint64) (rowsAffected int64, err error) {
	var (
		bucket model.BucketModel
		result *gorm.DB
	)

	result = db.Where("id = ?", id).Delete(&bucket)

	return result.RowsAffected, result.Error
}

// 恢复存储桶
func (b bucketRepo) Restore(db *gorm.DB, id uint64) (RowsAffected int64, err error) {
	var (
		bucket model.BucketModel
		result *gorm.DB
	)

	result = db.Unscoped().Where("id = ?", id).Find(&bucket).Update("deleted_at", gorm.Expr("NULL"))

	return result.RowsAffected, result.Error
}

// 强制删除存储桶
func (b bucketRepo) ForceDelete(db *gorm.DB, id uint64) (rowsAffected int64, err error) {
	var (
		bucket model.BucketModel
		result *gorm.DB
	)

	result = db.Unscoped().Where("id = ?", id).Delete(&bucket)

	return result.RowsAffected, result.Error
}

// 更新存储桶
func (b bucketRepo) Update(db *gorm.DB, id uint64, bucketMap map[string]interface{}) (err error) {
	var (
		bucket *model.BucketModel
	)
	// 检查账户是否存在
	if bucket, err = b.GetBucketByID(db, id); err != nil {
		return
	}

	return db.Model(bucket).Updates(bucketMap).Error
}

// 获取存储桶列表
func (b bucketRepo) GetBuckets(db *gorm.DB) (buckets []*model.BucketModel, err error) {
	var (
		result *gorm.DB
	)

	buckets = make([]*model.BucketModel, 0)
	result = db.Find(&buckets)

	return buckets, result.Error
}

// 通过ID获取存储桶
func (b bucketRepo) GetBucketByID(db *gorm.DB, id uint64) (bucket *model.BucketModel, err error) {
	var (
		result *gorm.DB
	)

	result = db.Where("id = ?", id).First(&bucket)
	err = result.Error

	return
}

// 通过名称获取存储桶
func (b bucketRepo) GetBucketByName(db *gorm.DB, name string) (bucket *model.BucketModel, err error) {
	var (
		result *gorm.DB
	)

	result = db.Where("name = ?", name).First(&bucket)
	err = result.Error

	return
}

// NewBucketRepo 实例化存储桶仓库
func NewBucketRepo() Repo {
	return &bucketRepo{}
}
