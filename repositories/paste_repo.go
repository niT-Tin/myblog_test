package repositories

import (
	"gorm.io/gorm"
	"myblog/models"
)

type IPaste interface {
	Insert(thing models.PasteThing) bool
	SelectSingle() (models.PasteThing, error)
	DeleteById(int64) bool
	DeleteAll() bool
}

type Paste struct {
	db *gorm.DB
}

func NewPaste(db *gorm.DB) IPaste {
	return &Paste{db: db}
}

// Insert 插入粘帖数据
func (p *Paste) Insert(thing models.PasteThing) bool {
	_, err2 := NoRepeat1(p.db.Create(&models.Paste{P: thing}), "插入内贴内容错误")
	if err2 != nil {
		return false
	}
	return true
}

// SelectSingle 获得单条粘帖数据
func (p *Paste) SelectSingle() (models.PasteThing, error) {
	var pst models.Paste
	_, err2 := NoRepeat1(p.db.First(&pst), "查找单条粘帖数据错误")
	return pst.P, err2
}

// DeleteById 根据id删除粘帖内容
func (p *Paste) DeleteById(pasteId int64) bool {
	_, err2 := NoRepeat1(p.db.Where("pasteId = ? ", pasteId).Delete(&models.Paste{}), "删除粘帖内容错误")
	if err2 != nil {
		return false
	}
	return true
}

// DeleteAll 删除所有粘帖内容
func (p *Paste) DeleteAll() bool {
	_, err2 := NoRepeat1(p.db.Delete(&models.Paste{}), "删除所有内容错误")
	if err2 != nil {
		return false
	}
	return true
}
