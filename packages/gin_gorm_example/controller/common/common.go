package common

import "github.com/gin-gonic/gin"

type Common struct {
}

func (c *Common) JsonSuccess(context *gin.Context, status int, h gin.H) {
	h["status"] = "success"
	context.JSON(status, h)
	return
}

func (c *Common) JsonFail(context *gin.Context, status int, message string) {
	context.JSON(status, gin.H{
		"status":  "fail",
		"message": message,
	})
}
