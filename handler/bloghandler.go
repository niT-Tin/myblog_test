package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myblog/initdb"
	"myblog/initdb/mysql"
	"myblog/models"
	"myblog/repositories"
	"myblog/service"
	"net/http"
	"strconv"
	"time"
)

// blog的 CRUD

var blogService service.IBlogService

func init() {
	blogService = service.NewBlogService(repositories.NewBlog(mysql.GetMYSqlConn()))
}

// GetBlog 获取博客数据
func GetBlog(c *gin.Context) {
	// 根据blogId 获取对应博客id并将字符串转换为int64类型
	blogId, err := strconv.ParseInt(c.Param("blogId"), 10, 10)
	// 如果转换存在错误则记录错误，并返回400响应码
	if err != nil {
		initdb.ErrorRecite(err, "get blog by id error", time.Now(), "handler getBlog")
		// 将错误提示信息序列化
		c.JSON(http.StatusBadRequest, models.Response{Code: http.StatusBadRequest, Message: "failed"})
		return
	}
	// 获取博客数据
	res := blogService.SelectByBlogId(repositories.Merge{BlogID: blogId})
	// 博客数据
	response := models.Response{Code: http.StatusOK, Message: "success", Data: res.B}
	// 返回成功信息
	c.JSON(http.StatusOK, response)
	return
}

func UpdateBlog(c *gin.Context) {

}

func DeleteBlog(c *gin.Context) {

}

func CreateBlogFromFile(c *gin.Context) {
	var m models.Blog
	err := c.ShouldBindJSON(&m)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", m)
}

func CreateBlog(c *gin.Context) {

}
