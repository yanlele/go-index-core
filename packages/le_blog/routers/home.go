package routers

import (
	"github.com/gin-gonic/gin"
	"le-blog/controllers"
	"le-blog/middleware"
)

func Home(r *gin.Engine) {
	home := r.Group("/")
	{
		// 首页
		home.GET("/", controllers.Index)

		article := home.Group("/article", middleware.Authorization)
		{
			//article.GET("/user", controllers.)
			article.GET("/user")
		}
	}
}
