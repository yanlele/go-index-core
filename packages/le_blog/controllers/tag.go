package controllers

import (
	"github.com/gin-gonic/gin"
	"le-blog/bootstrap/driver"
	"le-blog/modules"
	"le-blog/utils"
	"net/http"
)

func TagIndex(c *gin.Context) {
	var tags []modules.Tag

	err := driver.DB.Find(&tags).Error

	if err != nil {
		utils.RedirectBack(c)
	}

	auth := Auth{}.GetAuth(c)

	data := &struct {
		Auth
		Tags []modules.Tag
	}{auth, tags}

	c.HTML(http.StatusOK, "tagIndex", data)
}

// 标签
func GetArticleByTagName(c *gin.Context) {
	tagName := c.Param("id")
	if tagName == "" {
		utils.RedirectBack(c)
		return
	}

	var articles []modules.Article
	//dbQuery := driver.Db.Table("articles").Where("find_in_set('" + tagName + "', tags)").Order("view_num")
	dbQuery := driver.DB.Table("articles").Where("find_in_set('" + tagName + "', tags)").Order("view_num")

	paginate, err := (&modules.Pagination{}).Paginate(c, *dbQuery, &articles)

	if err != nil {
		panic(err)
	}

	auth := Auth{}.GetAuth(c)
	header := Header{"标签"}
	data := struct {
		Paginate modules.Pagination
		Auth
		Header
	}{*paginate, auth, header}
	c.HTML(http.StatusOK, "index", data)
}
