package models

import "time"

// Blog 博客结构
type Blog struct {
	ID            int64     `json:"blog_id" bson:"blog_id"`
	PosterID      int64     `json:"poster_id" bson:"poster_id"`
	CreatedAt     time.Time `json:"created_at" bson:"created_at"`
	LastUpdatedAt time.Time `json:"last_updated_at" bson:"last_updated_at"`
	Contents      Content   `json:"contents" bson:"contents"`
}
