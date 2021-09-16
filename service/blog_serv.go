package service

import (
	"errors"
	"myblog/initdb"
	"myblog/repositories"
	"time"
)

type IBlogService interface {
	InsertBlog(merge repositories.Merge) repositories.Merge
	UpdateBlog(repositories.Merge) repositories.Merge
	SelectByBlogId(repositories.Merge) repositories.Merge
	SelectAllBlog() repositories.Merge
	SelectByPosterId(repositories.Merge) repositories.Merge
	SelectAllByPosterId(merge repositories.Merge) repositories.Merge
	DeleteByBlogId(merge repositories.Merge) bool
	DeleteByPosterId(merge repositories.Merge) bool
	DeleteAllBlog() bool
}

func NoErrorRepeat(err error, msg string) repositories.Merge {
	// err may be useless
	var m repositories.Merge
	e := errors.New(msg)
	initdb.ErrorRecite(e, msg, time.Now(), "service")
	return m
}

type BlogService struct {
	BlogRepo repositories.IBlog
}

func NewBlogService(r repositories.IBlog) IBlogService {
	return &BlogService{BlogRepo: r}
}

// InsertBlog 插入博客服务
func (bs *BlogService) InsertBlog(merge repositories.Merge) repositories.Merge {
	insert, err := bs.BlogRepo.Insert(merge.B, merge.C)
	if err != nil {
		return NoErrorRepeat(err, "insert blog service error")
	}
	return repositories.Merge{RowsAffected: insert}
}

// UpdateBlog 更新博客服务
func (bs *BlogService) UpdateBlog(merge repositories.Merge) repositories.Merge {
	update, err := bs.BlogRepo.Update(merge.B, merge.C)
	if err != nil {
		return NoErrorRepeat(err, "update blog service error")
	}
	return repositories.Merge{RowsAffected: update}
}

// SelectByBlogId 根据博客id查询博客数据
func (bs *BlogService) SelectByBlogId(merge repositories.Merge) repositories.Merge {
	blog, content, err := bs.BlogRepo.SelectByBlogId(merge.BlogID)
	if err != nil {
		return NoErrorRepeat(err, "select by blog id error")
	}
	return repositories.Merge{
		B: blog,
		C: content,
	}
}

// SelectAllBlog 查询所有博客数据
func (bs *BlogService) SelectAllBlog() repositories.Merge {
	blogSlice, contentSlice, err := bs.BlogRepo.SelectAll()
	if err != nil {
		return NoErrorRepeat(err, "select all blog error")
	}
	return repositories.Merge{
		BS: blogSlice,
		CS: contentSlice,
	}
}

// SelectByPosterId 根据发布者id查询博客数据
func (bs *BlogService) SelectByPosterId(merge repositories.Merge) repositories.Merge {
	blog, content, err := bs.BlogRepo.SelectByPosterId(merge.PosterID, merge.BlogID)
	if err != nil {
		return NoErrorRepeat(err, "select by poster id error")
	}
	return repositories.Merge{
		B: blog,
		C: content,
	}
}

// SelectAllByPosterId 根据发布者id查询所有博客数据
func (bs *BlogService) SelectAllByPosterId(merge repositories.Merge) repositories.Merge {
	blogs, contents, err := bs.BlogRepo.SelectAllByPosterId(merge.PosterID)
	if err != nil {
		return NoErrorRepeat(err, "select all by poster id error")
	}
	return repositories.Merge{
		BS: blogs,
		CS: contents,
	}
}

// DeleteByBlogId 根据博客id删除博客数据
func (bs *BlogService) DeleteByBlogId(merge repositories.Merge) bool {
	return bs.BlogRepo.DeleteByBlogId(merge.BlogID)
}

// DeleteByPosterId 根据发布者id删除博客数据
func (bs *BlogService) DeleteByPosterId(merge repositories.Merge) bool {
	return bs.BlogRepo.DeleteByPosterId(merge.PosterID)
}

// DeleteAllBlog 删除所有博客数据
func (bs *BlogService) DeleteAllBlog() bool {
	return bs.BlogRepo.DeleteAll()
}
