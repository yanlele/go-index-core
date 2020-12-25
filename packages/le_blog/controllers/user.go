package controllers

import (
	"github.com/gin-gonic/gin"
	"le-blog/utils"
	"net/http"
)

// 表单提交数据
type formUser struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
}

// 注册页面
func Register(c *gin.Context) {
	auth := Auth{}.GetAuth(c)
	if auth.Id > 0 {
		utils.Redirect(c, "/")
	}

	data := struct {
		Title string
		Auth
	}{"注册", auth}

	c.HTML(http.StatusOK, "join", data)
}
