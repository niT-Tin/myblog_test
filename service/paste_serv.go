package service

import (
	"errors"
	"myblog/initdb"
	"myblog/models"
	"myblog/repositories"
	"time"
)

type IPasteService interface {
	Insert(thing string) bool
	SelectSingle() (string, error)
	DeletePasteById(int64) bool
	DeleteAllPaste() bool
}

func NoRepeatFromPasteService(isOk bool, msg string) bool {
	if !isOk {
		e := errors.New(msg)
		initdb.ErrorRecite(e, msg, time.Now(), "service")
		return false
	}
	return true
}

type PasteService struct {
	PasteRepo repositories.IPaste
}

func NewPasteService(r repositories.IPaste) IPasteService {
	return &PasteService{PasteRepo: r}
}

// Insert 插入粘帖内容
func (ps *PasteService) Insert(thing string) bool {
	return NoRepeatFromPasteService(ps.PasteRepo.Insert(models.PasteThing(thing)), "paste service insert 错误")
}

// SelectSingle 查询单条粘帖数据
func (ps *PasteService) SelectSingle() (string, error) {
	single, err := ps.PasteRepo.SelectSingle()
	if err != nil {
		e := errors.New("paste service selectSingle 错误")
		initdb.ErrorRecite(e, "paste service selectSingle 错误", time.Now(), "service")
		return "", e
	}
	return string(single), err
}

// DeletePasteById 根据粘帖数据id删除粘帖数据
func (ps *PasteService) DeletePasteById(id int64) bool {
	return NoRepeatFromPasteService(ps.PasteRepo.DeleteById(id), "paste service DeleteById 错误")
}

// DeleteAllPaste 删除所有粘帖数据
func (ps *PasteService) DeleteAllPaste() bool {
	return NoRepeatFromPasteService(ps.PasteRepo.DeleteAll(), "paste service Delete all 错误")
}
