package initRouter

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
简单的路由封装
*/
func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin")
	})
	return router
}
