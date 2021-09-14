package models

import "gorm.io/gorm"

// User 用户模型
type User struct {
	gorm.Model `redis:"-"`
	ID         int64  `json:"id" gorm:"column:id; autoIncrement;primaryKey"`
	NickName   string `json:"nickName" gorm:"column:nick_name"`
	UserName   string `json:"user_name" gorm:"column:user_name"`
	Password   string `json:"password" gorm:"column:password"`
}
