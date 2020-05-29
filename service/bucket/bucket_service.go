package bucket

import (
	"github.com/curder/go-qiniu-demo/model"
	"github.com/curder/go-qiniu-demo/repository/bucket"
)

// Service 用户服务接口定义
// 使用大写的service对外保留方法
type Service interface {
	Create(bucket model.BucketModel) (id uint64, err error)
	Delete(id uint64) (rowsAffected int64, err error)
	Restore(id uint64) (rowsAffected int64, err error)
	ForceDelete(id uint64) (rowsAffected int64, err error)
	Update(id uint64, bucketMap map[string]interface{}) error
	GetList() (buckets []*model.BucketModel, err error)
	GetByID(id uint64) (bucket *model.BucketModel, err error)
	GetByName(name string) (bucket *model.BucketModel, err error)
}

// BucketSvc 直接初始化，可以避免在使用时再实例化
var BucketSvc = NewBucketService()

// 用小写的 service 实现接口中定义的方法
type bucketSvc struct {
	bucketRepo bucket.Repo
}

// 创建存储桶
func (b *bucketSvc) Create(bucket model.BucketModel) (id uint64, err error) {
	if id, err = b.bucketRepo.Create(model.GetDB(), bucket); err != nil {
		return id, err
	}

	return id, nil
}

// 删除存储桶
func (b *bucketSvc) Delete(id uint64) (rowsAffected int64, err error) {
	if rowsAffected, err = b.bucketRepo.Delete(model.GetDB(), id); err != nil {
		return
	}

	return
}

// 恢复存储桶
func (b *bucketSvc) Restore(id uint64) (rowsAffected int64, err error) {
	if rowsAffected, err = b.bucketRepo.Restore(model.GetDB(), id); err != nil {
		return
	}

	return
}

// 强制删除存储桶
func (b *bucketSvc) ForceDelete(id uint64) (rowsAffected int64, err error) {
	if rowsAffected, err = b.bucketRepo.ForceDelete(model.GetDB(), id); err != nil {
		return
	}

	return
}

// 更新存储桶
func (b *bucketSvc) Update(id uint64, bucketMap map[string]interface{}) (err error) {
	if err = b.bucketRepo.Update(model.GetDB(), id, bucketMap); err != nil {
		return
	}
	return
}

// 获取存储桶列表
func (b *bucketSvc) GetList() (buckets []*model.BucketModel, err error) {
	if buckets, err = b.bucketRepo.GetList(model.GetDB()); err != nil {
		return
	}
	return
}

// 通过ID获取存储桶信息
func (b *bucketSvc) GetByID(id uint64) (bucket *model.BucketModel, err error) {
	if bucket, err = b.bucketRepo.GetByID(model.GetDB(), id); err != nil {
		return
	}
	return
}

// 通过名称获取存储桶
func (b *bucketSvc) GetByName(name string) (bucket *model.BucketModel, err error) {
	if bucket, err = b.bucketRepo.GetByName(model.GetDB(), name); err != nil {
		return
	}
	return
}

// NewBucketService 实例化一个bucketService
// 通过 NewService 函数初始化 Service 接口
// 依赖接口，不要依赖实现，面向接口编程
func NewBucketService() Service {
	return &bucketSvc{
		bucketRepo: bucket.NewBucketRepo(),
	}
}
