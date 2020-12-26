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
		user := home.Group("/user", middleware.Authorization)
		{
			user.GET("/update_pwd", controllers.UpdatePwd)
			user.POST("/update_pwd", controllers.DoUpdatePwd)
		}

		// 文章详情
		home.GET("/detail/:id", controllers.Detail)

		// 标签页面
		tag := home.Group("/tags")
		{
			tag.GET("/", controllers.TagIndex)
			tag.GET("/title/:name", controllers.GetArticleByTagName)
			tag.GET("/ajax/list", controllers.AjaxTags)
			tag.POST("/add", controllers.AddTags)
		}

		home.GET("/archives", controllers.Archives)

		// 注册
		home.GET("/join", controllers.Register)
		home.POST("/join", controllers.DoRegister)

		// sign in
		home.GET("/login", controllers.Login)
		home.POST("/login", controllers.DoLogin)
		home.GET("/logout", controllers.Logout)
	}
}
