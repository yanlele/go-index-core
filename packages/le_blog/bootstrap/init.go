package bootstrap

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"html/template"
	"le-blog/middleware"
	"le-blog/routers"
	"le-blog/utils"
	"net/http"
)

const CookieSessionKey = "blog_session"

func Init() *gin.Engine {
	app := gin.Default()

	// 完全不知道这个有啥用？
	//gob.Register(controllers.Auth{})

	// 添加cookie 和 session
	store := cookie.NewStore([]byte("secret"))
	app.Use(sessions.Sessions(CookieSessionKey, store))

	// 模板中添加函数
	app.SetFuncMap(template.FuncMap{
		"html":          utils.Html,
		"tagString2Map": utils.TagString2Map,
		"setLinkTitle":  utils.SetLinkTitle,
		"appUrl":        utils.AppUrl,
		"socialHtml":    utils.SocialHtml,
	})

	// 设置模板解析路径
	app.LoadHTMLGlob("./views/**/*")
	// 设置静态文件
	app.Static("/static", "./static")

	// 设置用户信息
	app.Use(middleware.SetAuth)
	app.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "404", nil)
	})

	routers.Api(app)
	routers.Home(app)

	return app
}
