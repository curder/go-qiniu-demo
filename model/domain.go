package model

import "time"

type DomainModel struct {
	ID uint64 `json:"id" gorm:"primary_key;AUTO_INCREMENT;column:id"`
	// Bucket     AccountModel `gorm:"foreignkey:BucketID"` // use BucketID as foreign key
	BucketID    uint64 `json:"bucket_id" gorm:"column:bucket_id"`
	Protocol    string `json:"protocol" gorm:"column:protocol"`
	Hostname    string `json:"hostname" gorm:"column:hostname"`
	Description string `json:"description" gorm:"description"`

	DeletedAt *time.Time `gorm:"column:deleted_at" json:"-"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"-"`
}

// TableName 表名
func (a *DomainModel) TableName() string {
	return "qiniu_domains"
}
