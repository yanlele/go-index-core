package test

import (
	"github.com/gin-gonic/gin"
	"go-index/packages/gin_series/06_upload_file/initRouter"
)

var router *gin.Engine

func init() {
	gin.SetMode(gin.TestMode)
	router = initRouter.SetupRouter()
}
