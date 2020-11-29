package v1

import (
	"gin-example/models"
	"gin-example/pkg/e"
	"gin-example/pkg/setting"
	"gin-example/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

func GetTags(context *gin.Context) {
	name := context.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	// 获取 url 的参数
	if arg := context.Query("state"); arg != "" {
		// 这个是类型转化， 当然也可以使用 strconv.Atoi 方法
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	// 分页处理
	data["lists"] = models.GetTags(util.GetPage(context), setting.PageSize, maps)

	// 获取总数
	data["total"] = models.GetTagTotal(maps)

	context.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": e.GetMsg(code),
		"data":    data,
	})
}

func AddTag(context *gin.Context) {

}

func EditTag(context *gin.Context) {

}

func DeleteTag(context *gin.Context) {

}
