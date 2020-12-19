package controllers

import (
	"github.com/gin-gonic/gin"
	"le-blog/bootstrap/driver"
	"le-blog/modules"
	"net/http"
)

type postArticle struct {
	Id             int    `form:"id" json:"id"`
	Title          string `form:"title" json:"title" binding:"required"`
	Tags           string `form:"tags" json:"tags" binding:"required"`
	Content        string `form:"content" json:"content" binding:"required"`
	EditorHtmlCode string `form:"editor-html-code" json:"editor-html-code"`
	DirectoryHtml  string `form:"directory_html" json:"directory_html"`
}

// article list 个人文章列表
func UserArticleList(c *gin.Context) {
	auth := Auth{}.GetAuth(c)
	var articles []modules.Article

	dbQuery := driver.DB.Table("articles").Where("user_id = ?", auth.Id).Order("created_at desc")
	paginate, err := (&modules.Pagination{}).Paginate(c, *dbQuery, &articles)
	if err != nil {
		panic(err)
	}

	data := struct {
		Paginate modules.Pagination
		Auth
	}{
		*paginate, auth,
	}

	c.HTML(http.StatusOK, "list", data)
}

// 创建文章页面
func CreateArticle(c *gin.Context) {
	auth := Auth{}.GetAuth(c)
	data := struct {
		Auth
	}{auth}
	c.HTML(http.StatusOK, "create-article", data)
}

// 保存文章接口
func SaveArticle(c *gin.Context) {
	var data postArticle
	if err := c.ShouldBind(&data); err != nil {

	}
}
