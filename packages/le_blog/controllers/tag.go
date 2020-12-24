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
