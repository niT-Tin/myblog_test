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

// Insert ??????????????????
func (b *Blog) Insert(blog *models.Blog, c *models.Content) (int64, error) {
	content := NewContent(b.db)
	_, err2 := content.Insert(c)
	if err2 != nil {
		t := time.Now()
		err10.ErrorRecite(err2, "????????????????????????", t, where)
		return 0, err2
	}
	repeatBlog := NoRepeatBlog(b.db.Create(blog), "??????????????????", ReturnTypeINT64, &Merge{})
	return repeatBlog.RowsAffected, repeatBlog.Error
}

// Update ??????????????????
func (b *Blog) Update(blog *models.Blog, m *models.Content) (int64, error) {
	blog.LastUpdatedAt = time.Now()
	_ = NoRepeatBlog(b.db.Model(m).Updates(m), "????????????????????????", ReturnTypeINT64, &Merge{})
	repeatBlog := NoRepeatBlog(b.db.Model(blog).Updates(blog), "??????????????????", ReturnTypeINT64, &Merge{})
	return repeatBlog.RowsAffected, repeatBlog.Error
}

// SelectByBlogId ????????????id??????????????????
func (b *Blog) SelectByBlogId(id int64) (*models.Blog, *models.Content, error) {
	c := models.Content{}
	m := Merge{
		B: &models.Blog{},
	}
	repeatBlog := NoRepeatBlog(b.db.First(m.B, id), "????????????????????????", ReturnTypeBlog, &m)
	contentId := m.B.ContentId // ????????????id
	row := b.db.Raw("select * from contents where id = ? ", contentId)
	row.Scan(&c)
	return repeatBlog.B, &c, repeatBlog.Error
}

// SelectAll ????????????????????????
func (b *Blog) SelectAll() ([]models.Blog, []models.Content, error) {
	m := &Merge{
		BS: []models.Blog{},
	}
	blog := NoRepeatBlog(b.db.Find(&m.BS), "????????????????????????", ReturnTypeSliceBlog, m)
	cs := []models.Content{}

	if len(m.BS) == 0 {
		err2 := errors.New("???????????????")
		t := time.Now()
		err10.ErrorRecite(err2, "???????????????", t, where)
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

// SelectByPosterId ???????????????id?????????id??????????????????
func (b *Blog) SelectByPosterId(postId, blogId int64) (*models.Blog, *models.Content, error) {
	m := &Merge{
		B: &models.Blog{},
	}
	blog := NoRepeatBlog(b.db.Where("poster_id = ? and id = ?", postId, blogId).First(m.B), "??????postId????????????", ReturnTypeBlog, m)
	var t models.Content
	b.db.Raw("select * from contents where id = ?", m.B.ContentId).Scan(&t)
	return blog.B, &t, blog.Error
}

// SelectAllByPosterId ???????????????id????????????????????????
func (b *Blog) SelectAllByPosterId(posterID int64) ([]models.Blog, []models.Content, error) {
	m := &Merge{
		BS: []models.Blog{},
	}
	blog := NoRepeatBlog(b.db.Where("poster_id = ?", posterID).Find(&m.BS), "????????????????????????", ReturnTypeSliceBlog, m)
	var cs []models.Content
	if len(m.BS) == 0 {
		err2 := errors.New("???????????????")
		t := time.Now()
		err10.ErrorRecite(err2, "???????????????", t, where)
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

// DeleteByBlogId ????????????id??????????????????
func (b *Blog) DeleteByBlogId(blogId int64) bool {
	m := &Merge{}
	NoRepeatBlog(b.db.Delete(&models.Blog{}, blogId), "??????????????????", ReturnTypeINT64, m)
	if m.RowsAffected != 0 {
		return true
	}
	return false
}

// DeleteByPosterId ???????????????id????????????
func (b *Blog) DeleteByPosterId(posterId int64) bool {
	m := &Merge{}
	b1 := models.Blog{}
	NoRepeatBlog(b.db.Where("poster_id = ?", posterId).Delete(&b1), "??????????????????", ReturnTypeINT64, m)
	if m.RowsAffected != 0 {
		return true
	}
	return false
}

// DeleteAll ??????????????????
func (b *Blog) DeleteAll() bool {
	tx := b.db.Delete(models.Blog{})
	if tx.RowsAffected == 0 {
		return false
	}
	return true
}
