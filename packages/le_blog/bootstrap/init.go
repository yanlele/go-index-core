package bootstrap

import (
	"github.com/gin-gonic/gin"
)

const CookieSessionKey = "blog_session"

func Init() *gin.Engine {
	app := gin.Default()
	return app
}
