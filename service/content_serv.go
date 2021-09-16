package service

import "myblog/repositories"

type IContentService interface {
	InsertContent(content repositories.Merge) repositories.Merge
	UpdateContent(content repositories.Merge) repositories.Merge
	SelectByContentId(idm repositories.Merge) repositories.Merge
	SelectAllContent() repositories.Merge
	DeleteByContentId(idm repositories.Merge) bool
	DeleteAllContent() bool
}

type ContentService struct {
	ContentRepo repositories.IContent
}

func NewContentService(c repositories.IContent) IContentService {
	return &ContentService{ContentRepo: c}
}

// InsertContent 插入内容
func (cs *ContentService) InsertContent(content repositories.Merge) repositories.Merge {
	insert, err := cs.ContentRepo.Insert(content.C)
	if err != nil {
		return NoErrorRepeat(err, "insert content error")
	}
	return repositories.Merge{
		RowsAffected: insert,
	}
}

// UpdateContent 更新内容
func (cs *ContentService) UpdateContent(content repositories.Merge) repositories.Merge {
	update, err := cs.ContentRepo.Update(content.C)
	if err != nil {
		return NoErrorRepeat(err, "update content error")
	}
	return repositories.Merge{
		RowsAffected: update,
	}
}

// SelectByContentId 根据内容id查询内容
func (cs *ContentService) SelectByContentId(idm repositories.Merge) repositories.Merge {
	content, err := cs.ContentRepo.SelectByContentId(idm.ContentID)
	if err != nil {
		return NoErrorRepeat(err, "select by content id error")
	}
	return repositories.Merge{
		C: content,
	}
}

// SelectAllContent 查询所有内容`
func (cs *ContentService) SelectAllContent() repositories.Merge {
	all, err := cs.ContentRepo.SelectAll()
	if err != nil {
		return NoErrorRepeat(err, "select all content error")
	}
	return repositories.Merge{
		CS: all,
	}
}

// DeleteByContentId 根据内容id删除内容
func (cs *ContentService) DeleteByContentId(idm repositories.Merge) bool {
	return cs.ContentRepo.DeleteByContentId(idm.ContentID)
}

// DeleteAllContent 删除所有内容
func (cs *ContentService) DeleteAllContent() bool {
	return cs.ContentRepo.DeleteAll()
}
