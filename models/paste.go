package models

import "gorm.io/gorm"

type PasteThing string

// Paste 黏贴的文件
type Paste struct {
	gorm.Model
	ID int64      `json:"id" gorm:"column:id"`
	P  PasteThing `json:"paste_thing" gorm:"column:paste_thing"`
}
