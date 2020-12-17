package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Api(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"code":    0,
				"message": "",
			})
		})
	}
}
