package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Redirect(c *gin.Context, location string) {
	c.Redirect(http.StatusFound, location)
	return
}
