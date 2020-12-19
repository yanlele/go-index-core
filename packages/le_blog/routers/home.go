package routers

import (
	"github.com/gin-gonic/gin"
	"le-blog/controllers"
)

func Home(r *gin.Engine) {
	home := r.Group("/")
	{
		// 首页
		home.GET("/", controllers.Index)
	}
}
