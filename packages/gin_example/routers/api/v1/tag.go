package v1

import (
	"gin-example/models"
	"gin-example/pkg/e"
	"gin-example/pkg/setting"
	"gin-example/pkg/util"
	"github.com/astaxie/beego/validation"
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
	name := context.Query("name")
	state := com.StrTo(context.DefaultQuery("state", "0")).MustInt()
	createdBy := context.Query("created_by")

	// 验证
	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistTagByName(name) {
			code = e.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": e.GetMsg(code),
		"data":    make(map[string]string),
	})
}

func EditTag(context *gin.Context) {
	id := com.StrTo(context.Query("id")).MustInt()
	name := context.Query("name")
	modifiedBy := context.Query("modified_by")

	valid := validation.Validation{}

	var state = -1
	if arg := context.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
	}
}

func DeleteTag(context *gin.Context) {

}
