package util

import (
	"gin-example/pkg/setting"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

/* 分页页码的获取方法 */
func GetPage(context *gin.Context) int {
	result := 0
	page, _ := com.StrTo(context.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}
	return result
}
