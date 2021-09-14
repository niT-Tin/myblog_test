package models

import "gorm.io/gorm"

// Content 内容结构
type Content struct {
	gorm.Model
	ID       int64  `json:"id" gorm:"column:id"`
	BlogId   int64  `json:"blog_id" gorm:"column:blog_id"`
	Text     string `json:"text" gorm:"column:text"`
	PicUrl   string `json:"pic_url" gorm:"column:pic_url"`
	VideoUrl string `json:"video_url" gorm:"column:video_url"`
}
