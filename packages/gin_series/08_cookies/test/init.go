package test

import (
	"github.com/gin-gonic/gin"
	"go-index/packages/gin_series/08_cookies/initRouter"
)

var router *gin.Engine

func init() {
	gin.SetMode(gin.TestMode)
	router = initRouter.SetupRouter()
}
