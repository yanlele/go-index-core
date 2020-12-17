package bootstrap

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const CookieSessionKey = "blog_session"

func Init() *gin.Engine {
	app := gin.Default()

	// 完全不知道这个有啥用？
	//gob.Register(controllers.Auth{})

	// 添加cookie 和 session
	store := cookie.NewStore([]byte("secret"))
	app.Use(sessions.Sessions(CookieSessionKey, store))


	return app
}
