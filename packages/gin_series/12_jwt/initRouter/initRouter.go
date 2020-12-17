package initRouter

import (
	"github.com/gin-gonic/gin"
	"go-index/packages/gin_series/12_jwt/handler/user"
	"go-index/packages/gin_series/12_jwt/middleware"
	"net/http"
	"time"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", middleware.Auth(), func(context *gin.Context) {
		context.JSON(http.StatusOK, time.Now().Unix())
	})
	router.GET("/login", user.CreateJwt)
	router.POST("/register", user.Register)
	return router
}
