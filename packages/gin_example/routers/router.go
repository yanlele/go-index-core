package routers

import (
	"gin-example/pkg/setting"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	gin.SetMode(setting.RunMode)
	router.GET("/test", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "test",
		})
	})

	return router
}
