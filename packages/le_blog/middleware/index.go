package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"le-blog/controllers"
	"le-blog/utils"
)

// 注入 auth
func SetAuth(c *gin.Context) {
	sess := sessions.Default(c)
	auth := sess.Get("auth")

	if auth != nil {
		c.Set("auth", auth)
	}
	c.Next()
}

func Authorization(c *gin.Context) {
	auth := controllers.Auth{}.GetAuth(c)
	if auth.Id == 0 {
		// 用户没有登录的情况
		utils.Redirect(c, "/login")
		return
	}
	c.Next()
}
