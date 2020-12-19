package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

type Response struct {
	Status uint
	Data   interface{}
	Msg    string
}

func (r *Response) SuccessResponse() gin.H {
	return gin.H{
		"status": 0,
		"data":   r.Data,
		"msg":    "",
	}
}

func (r *Response) FailResponse() gin.H {
	return gin.H{
		"status": r.Status,
		"data":   r.Data,
		"msg":    r.Msg,
	}
}

func Redirect(c *gin.Context, location string) {
	c.Redirect(http.StatusFound, location)
	return
}

// RedirectBack 重定向到上一次的页面
func RedirectBack(c *gin.Context)  {
	referer := c.GetHeader("Referer")
	pathInfo := ""
	if referer == "" {
		pathInfo = "/"
	} else {
		u, _ := url.Parse(referer)
		pathInfo = u.Path
	}

	c.Redirect(http.StatusFound, pathInfo)
	return
}

