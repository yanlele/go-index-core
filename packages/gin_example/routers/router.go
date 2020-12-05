package routers

import (
	"gin-example/middleware/jwt"
	"gin-example/pkg/setting"
	"gin-example/routers/api"
	v1 "gin-example/routers/api/v1"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	gin.SetMode(setting.RunMode)

	router.GET("/auth", api.GetAuth)

	router.GET("?swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := router.Group("/api/v1")
	apiv1.Use(jwt.Jwt())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
		// 获取指定id
		apiv1.GET("/tags/:id", v1.GetOneTag)

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return router
}
