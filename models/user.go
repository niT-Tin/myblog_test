package models

import "gorm.io/gorm"

// User 用户模型
type User struct {
	gorm.Model
	ID       int64  `json:"id" gorm:"id"`
	UserName string `json:"user_name" gorm:"user_name"`
	Password string `json:"password" gorm:"password"`
}
