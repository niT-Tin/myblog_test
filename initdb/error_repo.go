package initdb

import (
	"gorm.io/gorm"
	"myblog/initdb/mysql"
	"myblog/models"
	"time"
)

type IError interface {
	Insert(error2 *models.Error) bool
	Select() ([]models.Error, error)
}

type Error struct {
	Db *gorm.DB
}

func NewError(db *gorm.DB) IError {
	return &Error{Db: db}
}

// Insert 插入错误数据
func (e *Error) Insert(error2 *models.Error) bool {
	return e.Db.Create(error2).RowsAffected != 0
}

// Select 查询错误数据
func (e *Error) Select() ([]models.Error, error) {
	var es []models.Error
	find := e.Db.Find(&es)
	if find.Error != nil {
		t := time.Now()
		e.Insert(&models.Error{
			Message: "查询错误数据表出现错误",
			Where:   "error repo",
			When:    t,
		})
	}
	return es, find.Error
}

// ErrorRecite 插入错误数据
func ErrorRecite(e error, msg string, when time.Time, where string) {
	newError := NewError(mysql.DB)
	newError.Insert(&models.Error{
		Message: msg,
		When:    when,
		Where:   where,
	})
}

// ErrorRetriever 查询错误数据
func ErrorRetriever() (errs []models.Error) {
	newError := NewError(mysql.DB)
	errs, err1 := newError.Select()
	if err1 != nil {
		t := time.Now()
		ErrorRecite(err1, "获取错误信息时发生错误", t, "ErrorRetriever")
		return
	}
	return errs
}
