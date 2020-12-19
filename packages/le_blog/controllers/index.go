package controllers

import (
	"github.com/gin-gonic/gin"
	"le-blog/bootstrap/driver"
	"le-blog/modules"
)

// 项目首页
func Index(c *gin.Context) {
	var articles []modules.Article
	dbQuery := driver.DB.Table("articles").Order("created_at desc")

}

