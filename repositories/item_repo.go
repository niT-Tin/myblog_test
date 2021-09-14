package repositories

import (
	"gorm.io/gorm"
	"myblog/models"
)

type IItem interface {
	SelectAll() ([]models.Item, error)
	SelectByLimit(int) ([]models.Item, error)
}

type ItemStruct struct {
	db *gorm.DB
}

func NewItemStruct(db *gorm.DB) IItem {
	return &ItemStruct{
		db: db,
	}
}

// SelectAll 查询所有item
func (i *ItemStruct) SelectAll() ([]models.Item, error) {
	var is []models.Item
	_, err2 := NoRepeat1(i.db.Find(&is), "查询item信息错误")
	return is, err2
}

// SelectByLimit 根据limit查询item
func (i *ItemStruct) SelectByLimit(limit int) ([]models.Item, error) {
	var is []models.Item
	_, err2 := NoRepeat1(i.db.Limit(limit).Find(&is), "limit查询错误")
	return is, err2
}
