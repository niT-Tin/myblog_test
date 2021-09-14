package repositories

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	err2 "myblog/initdb"
	"myblog/models"
	"time"
)

type IContent interface {
	Insert(content *models.Content) (int64, error)
	Update(content *models.Content) (int64, error)
	SelectByContentId(int64) (*models.Content, error)
	SelectAll() ([]models.Content, error)
	DeleteByContentId(int64) bool
	DeleteAll() bool
}

type Content struct {
	db *gorm.DB
}

func NewContent(db *gorm.DB) IContent {
	return &Content{db: db}
}

func NoRepeat1(db *gorm.DB, msg string) (int64, error) {
	if db.RowsAffected == 0 {
		err1 := errors.New(msg)
		t := time.Now()
		err2.ErrorRecite(err1, msg, t, where)
		return 0, err1
	}
	return db.RowsAffected, nil
}

// Insert 插入内容数据
func (c *Content) Insert(content *models.Content) (int64, error) {
	return NoRepeat1(c.db.Create(content), "插入博客内容错误")
}

// Update 更新博客内容
func (c *Content) Update(content *models.Content) (int64, error) {
	return NoRepeat1(c.db.Model(content).Updates(content), "更新博客内容错误")
}

// SelectByContentId 根据内容id查询内容
func (c *Content) SelectByContentId(contentId int64) (*models.Content, error) {
	var content models.Content
	_, err2 := NoRepeat1(c.db.Where("contentId = ? ", contentId).First(&content), "查询博客内容错误")
	return &content, err2
}

// SelectAll 查询所有内容数据
func (c *Content) SelectAll() ([]models.Content, error) {
	var cs []models.Content
	_, err2 := NoRepeat1(c.db.Find(&cs), "查询所有内容错误")
	return cs, err2
}

// DeleteByContentId 根据内容id删除数据
func (c *Content) DeleteByContentId(contentId int64) bool {
	_, err2 := NoRepeat1(c.db.Where("contentId = ?", contentId).Delete(models.Content{}), "删除博客内容错误")
	if err2 != nil {
		return false
	}
	return true
}

// DeleteAll 删除所有内容数据
func (c *Content) DeleteAll() bool {
	_, err2 := NoRepeat1(c.db.Delete(&models.Content{}), "删除所有内容错误")
	if err2 != nil {
		return false
	}
	return true
}
