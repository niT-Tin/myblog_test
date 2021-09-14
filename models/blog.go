package models

import (
	"gorm.io/gorm"
	"time"
)

// Blog 博客结构
type Blog struct {
	gorm.Model
	ID            int64     `json:"blog_id" bson:"blog_id" gorm:"autoIncrement;primaryKey; column:id"`
	PosterID      int64     `json:"poster_id" gorm:"column:poster_id"`
	CreatedAt     time.Time `json:"created_at" gorm:"column:created_at"`
	LastUpdatedAt time.Time `json:"last_updated_at" gorm:"column:last_updated_at"`
	ContentId     int64     `json:"content_id" gorm:"column:content_id"`
}
