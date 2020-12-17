package initRouter

import (
	"github.com/gin-gonic/gin"
	"go-index/packages/gin_series/08_cookies/handler"
	"go-index/packages/gin_series/08_cookies/middleware"
	"go-index/packages/gin_series/08_cookies/utils"
	"net/http"
	"path/filepath"
)

func SetupRouter() *gin.Engine {
	router := gin.New()
	router.Use(middleware.Logger(), gin.Recovery())
	if mode := gin.Mode(); mode == gin.TestMode {
		router.LoadHTMLGlob("../templates/*")
	} else {
		router.LoadHTMLGlob("packages/gin_series/08_cookies/templates/*")
	}
	router.StaticFile("/favicon.ico", "packages/gin_series/statics/favicon.ico")
	router.Static("/statics", "packages/gin_series/statics")

	// 添加头像静态文件的位置
	router.StaticFS("/avatar", http.Dir(filepath.Join(utils.RootPath(), "avatar")))
	index := router.Group("/")
	{
		index.Any("", handler.Index)
	}
	// 添加 user
	userRouter := router.Group("/user")
	{
		userRouter.POST("/register", handler.UserRegister)
		userRouter.POST("/login", handler.UserLogin)
		userRouter.GET("/profile/", middleware.Auth(), handler.UserProFile)
		userRouter.POST("/update", middleware.Auth(), handler.UpdateUserProfile)
	}

	return router
}
