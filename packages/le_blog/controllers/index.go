package controllers

import (
	"github.com/gin-gonic/gin"
	"le-blog/bootstrap/driver"
	"le-blog/modules"
	"net/http"
)

// 项目首页
func Index(c *gin.Context) {
	var articles []modules.Article
	dbQuery := driver.DB.Table("articles").Order("created_at desc")
	paginate, err := (&modules.Pagination{}).Paginate(c, *dbQuery, &articles)
	if err != nil {
		panic(err)
	}
	auth := Auth{}.GetAuth(c)
	header := Header{""}
	data := struct {
		Paginate modules.Pagination
		Auth
		Header
	}{
		*paginate, auth, header,
	}

	c.HTML(http.StatusOK, "index", data)
}
