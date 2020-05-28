package model

import (
	"time"
)

// UserModel User represents a registered user.
type UserModel struct {
	ID       uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	Username string `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=4,max=128" json:"-"`
	//DeletedAt time.Time `gorm:"column:deleted_at" json:"-"`
	CreatedAt time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"-"`
}

// UserInfo 对外暴露的结构体
type UserInfo struct {
	ID        uint64 `json:"id" example:"1"`
	Username  string `json:"username" example:"张三"`
	Email     string `json:"email" example:"example@exmaple.com"`
	CreatedAt string `json:"created_at" example:"2020-03-23 20:00:00"`
	UpdatedAt string `json:"updated_at" example:"2020-03-23 20:00:00"`
}

// TableName 表名
func (u *UserModel) TableName() string {
	return "users"
}

// Token represents a JSON web token.
type Token struct {
	Token string `json:"token"`
}
