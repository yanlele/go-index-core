package initRouter

import (
	"github.com/gin-gonic/gin"
	"go-index/packages/gin_series/02_router/handler"
	"net/http"
	"strings"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	index := router.Group("/")
	{
		index.Any("", retHelloGinAndMethod)
	}

	userRouter := router.Group("/user")
	{
		userRouter.GET("/:name", handler.UserSave)
		userRouter.GET("/", handler.UserSaveByQuery)
	}
	return router
}

func retHelloGinAndMethod(context *gin.Context) {
	context.String(http.StatusOK, "hello gin "+strings.ToLower(context.Request.Method)+" method")
}
