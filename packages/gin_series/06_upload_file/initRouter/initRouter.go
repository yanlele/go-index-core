package initRouter

import (
	"github.com/gin-gonic/gin"
	"go-index/packages/gin_series/06_upload_file/handler"
	"go-index/packages/gin_series/06_upload_file/utils"
	"net/http"
	"path/filepath"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	if mode := gin.Mode(); mode == gin.TestMode {
		router.LoadHTMLGlob("../templates/*")
	} else {
		router.LoadHTMLGlob("packages/gin_series/06_upload_file/templates/*")
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
		userRouter.GET("/profile/", handler.UserProFile)
		userRouter.POST("/update", handler.UpdateUserProfile)
	}

	return router
}
