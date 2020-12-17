package initRouter

import (
	"github.com/gin-gonic/gin"
	"go-index/packages/gin_series/05_db/handler"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	if mode := gin.Mode(); mode == gin.TestMode {
		router.LoadHTMLGlob("../templates/*")
	} else {
		router.LoadHTMLGlob("packages/gin_series/05_db/templates/*")
	}
	router.StaticFile("/favicon.ico", "packages/gin_series/statics/favicon.ico")
	router.Static("/statics", "packages/gin_series/statics")
	index := router.Group("/")
	{
		index.Any("", handler.Index)
	}
	// 添加 user
	userRouter := router.Group("/user")
	{
		userRouter.POST("/register", handler.UserRegister)
		userRouter.POST("/login", handler.UserLogin)
	}

	return router
}
