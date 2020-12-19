package controllers

import (
	"github.com/gin-gonic/gin"
	"le-blog/bootstrap/driver"
	"le-blog/modules"
	"le-blog/utils"
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
		response := utils.Response{
			Status: 403,
			Data:   nil,
			Msg:    err.Error(),
		}

		c.JSON(http.StatusBadRequest, response.FailResponse())
		return
	}

	auth := Auth{}.GetAuth(c)

	// 计算文章简介
	introduction := utils.Html2Str(data.EditorHtmlCode)

	if len(introduction) > 100 {
		introduction = string(([]rune(introduction))[:100])
	} else {
		introduction = string(([]rune(introduction))[:len(introduction)])
	}

	article := modules.Article{
		Title:         data.Title,
		Introduction:  introduction,
		ContentMd:     data.Content,
		UserID:        auth.Id,
		Tags:          data.Tags,
		ContentHtml:   data.EditorHtmlCode,
		DirectoryHtml: data.DirectoryHtml,
	}

	var err error
	if data.Id == 0 {
		// 保存数据
		err = driver.DB.Create(&article).Error

		// todo 启动协程
	} else {
		article.ID = uint(data.Id)
		// 更新数据
		err = driver.DB.Save(&article).Error
	}

	if err != nil {
		response := utils.Response{
			Status: 500,
			Data: nil,
			Msg: err.Error(),
		}

		c.JSON(http.StatusOK, response.FailResponse())
		return
	}

	// todo 处理文章的tags , 启动协程

	response:= utils.Response{
		Status: 0,
		Data: article,
		Msg: "",
	}

	c.JSON(http.StatusOK, response.SuccessResponse())
	return
}
