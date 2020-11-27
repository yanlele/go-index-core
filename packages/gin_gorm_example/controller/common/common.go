package common

import "github.com/gin-gonic/gin"

type Common struct {
}

func (Common) JsonSuccess(context *gin.Context, status int, h gin.H) {
	h["status"] = "success"
	context.JSON(status, h)
	return
}

func (Common) JsonFail(context *gin.Context, status int, message string) {
	context.JSON(status, gin.H{
		"status":  "fail",
		"message": message,
	})
}
