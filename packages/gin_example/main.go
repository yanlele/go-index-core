package main

import (
	"gin-example/pkg/setting"
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	router := gin.Default()
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
	var port = ":" + strconv.Itoa(setting.HTTPPort)
	router.Run(port)
}
