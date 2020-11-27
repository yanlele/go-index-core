package response

import "github.com/gin-gonic/gin"

type CommonResponse struct {
}

func (c *CommonResponse) JsonSuccess(context *gin.Context, status int, h gin.H) {
	h["status"] = "success"
	context.JSON(status, h)
	return
}

func (c *CommonResponse) JsonFail(context *gin.Context, status int, message string) {
	context.JSON(status, gin.H{
		"status":  "fail",
		"message": message,
	})
}
