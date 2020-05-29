package model

import "time"

// 存储桶模型
type BucketModel struct {
	ID          uint64       `json:"id" gorm:"primary_key;AUTO_INCREMENT;column:id"`
	// Account     AccountModel `gorm:"foreignkey:AccountID"` // use AccountID as foreign key
	AccountID   uint64       `json:"account_id" gorm:"column:account_id"`
	Name        string       `json:"name" gorm:"column:name"`
	Description string       `json:"description" gorm:"description"`

	DeletedAt *time.Time `gorm:"column:deleted_at" json:"-"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"-"`
}

// TableName 表名
func (a *BucketModel) TableName() string {
	return "qiniu_buckets"
}
