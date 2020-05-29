package model

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type AccountModel struct {
	ID          uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	Email       string `json:"email" gorm:"column:email"`
	AccessKey   string `json:"access_key" gorm:"access_key"`
	SecretKey   string `json:"secret_key" gorm:"secret_key"`
	Description string `json:"description" gorm:"description"`

	DeletedAt *time.Time `gorm:"column:deleted_at" json:"-"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"-"`
}

// TableName 表名
func (a *AccountModel) TableName() string {
	return "qiniu_accounts"
}

// Validate the fields.
func (a *AccountModel) Validate() error {
	var (
		validate *validator.Validate
	)
	validate = validator.New()

	return validate.Struct(a)
}
