package models

// Content 内容结构
type Content struct {
	Text     string   `json:"text" bson:"text"`
	PicUrl   []string `json:"pic_url" bson:"pic_url"`
	VideoUrl []string `json:"video_url" bson:"video_"`
}
