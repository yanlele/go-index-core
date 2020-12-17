package test

import (
	"github.com/gin-gonic/gin"
	"go-index/packages/gin_series/04_form/initRouter"
)

var router *gin.Engine

func init() {
	gin.SetMode(gin.TestMode)
	router = initRouter.SetupRouter()
}
