package repositories

import (
	"errors"
	"gorm.io/gorm"
	err10 "myblog/initdb"
	"myblog/models"
	"time"
)

var (
	where = "blog_repo"
	err   error
)

const (
	ReturnTypeINT64 = iota
	ReturnTypeBlog
	ReturnTypeSliceBlog
)

type IBlog interface {
	Insert(blog *models.Blog, m *models.Content) (int64, error)
	Update(blog *models.Blog, m *models.Content) (int64, error)
	SelectByBlogId(int64) (*models.Blog, *models.Content, error)
	SelectAll() ([]models.Blog, []models.Content, error)
	SelectByPosterId(int64, int64) (*models.Blog, *models.Content, error)
	SelectAllByPosterId(int64) ([]models.Blog, []models.Content, error)
	DeleteByBlogId(int64) bool
	DeleteByPosterId(int64) bool
	DeleteAll() bool
}

type Merge struct {
	B            *models.Blog     `json:"blog,omitempty"`
	BS           []models.Blog    `json:"blog_slice,omitempty"`
	C            *models.Content  `json:"content,omitempty"`
	CS           []models.Content `json:"content_slice,omitempty"`
	IS           []models.Item    `json:"item_slice,omitempty"`
	U            models.User      `json:"user,omitempty"`
	US           []models.User    `json:"user_slice,omitempty"`
	IsSuccess    bool             `json:"is_success,omitempty"`
	Token        string           `json:"token,omitempty"`
	ContentID    int64            `json:"content_id,omitempty"`
	PosterID     int64            `json:"poster_id"`
	BlogID       int64            `json:"blog_id"`
	RowsAffected int64            `json:"rows_affected"`
	Error        error            `json:"error"`
}

type Blog struct {
	db *gorm.DB
}

func NewBlog(db *gorm.DB) IBlog {
	return &Blog{db: db}
}

func NoRepeatBlog(db *gorm.DB, ErrMsg string, returnType int, merge *Merge) *Merge {
	switch returnType {
	case ReturnTypeINT64:
		if db.RowsAffected == 0 {
			err = errors.New(ErrMsg)
			t := time.Now()
			err10.ErrorRecite(err, ErrMsg, t, where)
			return &Merge{
				RowsAffected: 0,
				Error:        err,
			}
		}
		return &Merge{
			RowsAffected: db.RowsAffected,
			Error:        nil,
		}
	case ReturnTypeBlog:
		if db.RowsAffected == 0 {
			err = errors.New(ErrMsg)
			t := time.Now()
			err10.ErrorRecite(err, ErrMsg, t, where)
			return &Merge{
				B:     &models.Blog{},
				Error: err,
			}
		}
		return merge
	case ReturnTypeSliceBlog:
		if db.RowsAffected == 0 {
			err = errors.New(ErrMsg)
			t := time.Now()
			err10.ErrorRecite(err, ErrMsg, t, where)
			return &Merge{
				BS:    []models.Blog{},
				Error: err,
			}
		}
		return merge
	default:
		return &Merge{}
	}
}

// Insert 插入博客数据
func (b *Blog) Insert(blog *models.Blog, c *models.Content) (int64, error) {
	content := NewContent(b.db)
	_, err2 := content.Insert(c)
	if err2 != nil {
		t := time.Now()
		err10.ErrorRecite(err2, "插入博客内容错误", t, where)
		return 0, err2
	}
	repeatBlog := NoRepeatBlog(b.db.Create(blog), "插入博客错误", ReturnTypeINT64, &Merge{})
	return repeatBlog.RowsAffected, repeatBlog.Error
}

// Update 更新博客数据
func (b *Blog) Update(blog *models.Blog, m *models.Content) (int64, error) {
	blog.LastUpdatedAt = time.Now()
	_ = NoRepeatBlog(b.db.Model(m).Updates(m), "更新博客内容错误", ReturnTypeINT64, &Merge{})
	repeatBlog := NoRepeatBlog(b.db.Model(blog).Updates(blog), "更新博客错误", ReturnTypeINT64, &Merge{})
	return repeatBlog.RowsAffected, repeatBlog.Error
}

// SelectByBlogId 通过博客id查询博客数据
func (b *Blog) SelectByBlogId(id int64) (*models.Blog, *models.Content, error) {
	c := models.Content{}
	m := Merge{
		B: &models.Blog{},
	}
	repeatBlog := NoRepeatBlog(b.db.First(m.B, id), "查询错误数据错误", ReturnTypeBlog, &m)
	contentId := m.B.ContentId // 获取内容id
	row := b.db.Raw("select * from contents where id = ? ", contentId)
	row.Scan(&c)
	return repeatBlog.B, &c, repeatBlog.Error
}

// SelectAll 查询所有博客数据
func (b *Blog) SelectAll() ([]models.Blog, []models.Content, error) {
	m := &Merge{
		BS: []models.Blog{},
	}
	blog := NoRepeatBlog(b.db.Find(&m.BS), "查询全部博客错误", ReturnTypeSliceBlog, m)
	cs := []models.Content{}

	if len(m.BS) == 0 {
		err2 := errors.New("无博客内容")
		t := time.Now()
		err10.ErrorRecite(err2, "无博客内容", t, where)
		return []models.Blog{}, []models.Content{}, err2
	}
	for i := 0; i < len(m.BS); i++ {
		row := b.db.Raw("select * from contents where id = ?", m.BS[i].ContentId)
		var t models.Content
		row.Scan(&t)
		cs = append(cs, t)
	}
	return blog.BS, cs, blog.Error
}

// SelectByPosterId 根据发布者id和博客id查询博客数据
func (b *Blog) SelectByPosterId(postId, blogId int64) (*models.Blog, *models.Content, error) {
	m := &Merge{
		B: &models.Blog{},
	}
	blog := NoRepeatBlog(b.db.Where("poster_id = ? and id = ?", postId, blogId).First(m.B), "通过postId查询错误", ReturnTypeBlog, m)
	var t models.Content
	b.db.Raw("select * from contents where id = ?", m.B.ContentId).Scan(&t)
	return blog.B, &t, blog.Error
}

// SelectAllByPosterId 根据发布者id查询所有博客数据
func (b *Blog) SelectAllByPosterId(posterID int64) ([]models.Blog, []models.Content, error) {
	m := &Merge{
		BS: []models.Blog{},
	}
	blog := NoRepeatBlog(b.db.Where("poster_id = ?", posterID).Find(&m.BS), "查询所有博客错误", ReturnTypeSliceBlog, m)
	var cs []models.Content
	if len(m.BS) == 0 {
		err2 := errors.New("无博客内容")
		t := time.Now()
		err10.ErrorRecite(err2, "无博客内容", t, where)
		return []models.Blog{}, []models.Content{}, err2
	}
	for i := 0; i < len(m.BS); i++ {
		row := b.db.Raw("select * from contents where id = ?", m.BS[i].ContentId)
		var t models.Content
		row.Scan(&t)
		cs = append(cs, t)
	}
	return blog.BS, cs, blog.Error
}

// DeleteByBlogId 根据博客id删除博客内容
func (b *Blog) DeleteByBlogId(blogId int64) bool {
	m := &Merge{}
	NoRepeatBlog(b.db.Delete(&models.Blog{}, blogId), "删除博客错误", ReturnTypeINT64, m)
	if m.RowsAffected != 0 {
		return true
	}
	return false
}

// DeleteByPosterId 根据发布者id删除博客
func (b *Blog) DeleteByPosterId(posterId int64) bool {
	m := &Merge{}
	b1 := models.Blog{}
	NoRepeatBlog(b.db.Where("poster_id = ?", posterId).Delete(&b1), "删除博客错误", ReturnTypeINT64, m)
	if m.RowsAffected != 0 {
		return true
	}
	return false
}

// DeleteAll 删除所有博客
func (b *Blog) DeleteAll() bool {
	tx := b.db.Delete(models.Blog{})
	if tx.RowsAffected == 0 {
		return false
	}
	return true
}
