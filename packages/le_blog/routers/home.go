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
			article.GET("/user", controllers.UserArticleList)
			article.GET("/create", controllers.CreateArticle)
			article.POST("/create", controllers.SaveArticle)
			article.GET("/edit/:id", controllers.EditArticle)
			article.GET("/delete/:id", controllers.DelArticle)
		}

		// 个人中心
		//user := home.Group("/user", middleware.Authorization)
		//{
		//	user.GET("/update_pwd", controllers)
		//}
	}
}
