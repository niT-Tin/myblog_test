package repositories

import (
	"gorm.io/gorm"
	"myblog/models"
	"time"
)

type IError interface {
	Insert(error2 *models.Error) bool
	Select(error2 *models.Error) ([]models.Error, error)
}

type Error struct {
	db *gorm.DB
}

func NewError(db *gorm.DB) IError {
	return &Error{db: db}
}

// Insert 插入错误数据
func (e *Error) Insert(error2 *models.Error) bool {
	return e.db.Create(error2).RowsAffected != 0
}

// Select 查询错误数据
func (e *Error) Select(error2 *models.Error) ([]models.Error, error) {
	var es []models.Error
	find := e.db.Find(&es)
	if find.Error != nil {
		e.Insert(&models.Error{
			E:       find.Error,
			Message: "查询错误数据表出现错误",
			Where:   "error repo",
			When:    time.Now(),
		})
	}
	return es, find.Error
}
