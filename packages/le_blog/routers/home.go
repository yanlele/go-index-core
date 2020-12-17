package routers

import "github.com/gin-gonic/gin"

func Home(r *gin.Engine) {
	home := r.Group("/")
	{
		home.GET("/")
	}
}
