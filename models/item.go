package models

import (
	"gorm.io/gorm"
	"time"
)

type Item struct {
	gorm.Model
	PicUrl   string    `json:"item_pic_url" gorm:"column:item_pic_url"`
	Text     string    `json:"item_text" gorm:"column:item_text"`
	PostedAt time.Time `json:"item_posted_at" gorm:"column:item_posted_at"`
}
