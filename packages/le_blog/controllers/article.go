package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"le-blog/bootstrap/driver"
	"le-blog/modules"
	"le-blog/services"
	"le-blog/utils"
	"net/http"
	"strconv"
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
		go services.SetArticleArchive(&article)
	} else {
		article.ID = uint(data.Id)
		// 更新数据
		err = driver.DB.Save(&article).Error
	}

	if err != nil {
		response := utils.Response{
			Status: 500,
			Data:   nil,
			Msg:    err.Error(),
		}

		c.JSON(http.StatusOK, response.FailResponse())
		return
	}

	// todo 处理文章的tags , 启动协程
	go services.HandleTags(data.Tags)

	response := utils.Response{
		Status: 0,
		Data:   article,
		Msg:    "",
	}

	c.JSON(http.StatusOK, response.SuccessResponse())
	return
}

// detail 文章详情
func Detail(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var article modules.Article
	article.ID = uint(id)

	err := driver.DB.First(&article).Error
	if err != nil {
		panic(err)
	}

	// article 的流量 + 1
	article.ViewNum = article.ViewNum + 1
	driver.DB.Save(&article)

	auth := Auth{}.GetAuth(c)
	header := Header{Title: article.Title}
	data := &struct {
		Article modules.Article
		Auth
		Header
	}{
		article,
		auth,
		header,
	}
	c.HTML(http.StatusOK, "detail", data)
}

// EditArticle 编辑文章
func EditArticle(c *gin.Context) {
	id, _ := c.Params.Get("id")
	auth := (&Auth{}).GetAuth(c)
	var article modules.Article

	err := driver.DB.Where("id = ?", id).First(&article).Error

	if err != nil {
		panic(err)
	}

	data := struct {
		modules.Article
		Auth
	}{article, auth}

	c.HTML(http.StatusOK, "edit-article", data)
}

// DelArticle 删除文章
func DelArticle(c *gin.Context) {
	id := c.Param("id")
	auth := Auth{}.GetAuth(c)

	article := modules.Article{
		Model:        gorm.Model{},
		Title:        "",
		Introduction: "",
		ContentMd:    "",
		ContentHtml:  "",
		UserID:       0,
		User:         modules.User{},
		Tags:         "",
		ViewNum:      0,
	}

	err := driver.DB.Where("id = ？and user_id = ?", id, auth.Id).Find(&article).Error
	if err != nil {
		utils.RedirectBack(c)
	}

	if article.ID == 0 {
		utils.RedirectBack(c)
	}

	err = driver.DB.Delete(&article).Error
	if err != nil {
		// session 写入 error
	}

	utils.RedirectBack(c)
}
