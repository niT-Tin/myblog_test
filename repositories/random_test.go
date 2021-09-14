package repositories

import (
	"fmt"
	"myblog/initdb/mysql"
	"myblog/models"
	"testing"
	"time"
)

func TestBlog_Insert(t *testing.T) {
	c := models.Content{
		BlogId:   1,
		Text:     "!242325",
		PicUrl:   "http://234234.com",
		VideoUrl: "http://234234.com",
	}
	b := models.Blog{
		PosterID:      1,
		CreatedAt:     time.Now(),
		LastUpdatedAt: time.Now(),
		ContentId:     1,
	}
	blog := NewBlog(mysql.GetMYSqlConn())
	insert, err2 := blog.Insert(&b, &c)
	if insert == 0 || err2 != nil {
		t.Errorf("插入数据错误")
	}
}

func TestBlog_SelectAllByPosterId(t *testing.T) {
	var id int64 = 1
	blog := NewBlog(mysql.GetMYSqlConn())
	posterId, contents, err2 := blog.SelectAllByPosterId(id)
	if err2 != nil {
		t.Errorf("查询数据错误")
	} else {
		fmt.Printf("blog := %+v\n", posterId)
		fmt.Printf("contents := %+v\n", contents)
	}
}
