package models

import (
	"gorm.io/gorm"
	"time"
)

// Error 错误结构体
type Error struct {
	gorm.Model
	//E       error     `json:"error" gorm:"embedded"`
	Message string    `json:"message" gorm:"column:message"`
	Where   string    `json:"where" gorm:"column:where"`
	When    time.Time `json:"when" gorm:"embedded"`
}
