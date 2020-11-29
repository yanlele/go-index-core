package v1

import (
	"gin-example/models"
	"gin-example/pkg/e"
	"gin-example/pkg/setting"
	"gin-example/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"log"
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

/* 编辑tag */
func EditTag(context *gin.Context) {
	// 获取参数
	id := com.StrTo(context.Param("id")).MustInt()
	name := context.Query("name")
	modifiedBy := context.Query("modified_by")

	// 初始化验证
	valid := validation.Validation{}

	// 验证状态范围
	var state = -1
	if arg := context.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	// 验证状态
	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")

	code := e.INVALID_PARAMS

	// 没有错误
	if !valid.HasErrors() {
		code = e.SUCCESS
		// 判定是否有这个 tag 存在
		if models.ExistTagById(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}
			models.EditTag(id, data)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	} else {
		log.Panicln("has error", valid.Errors)
	}

	context.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": e.GetMsg(code),
		"data":    make(map[string]interface{}),
	})
}

/* 根据 id 删除tag */
func DeleteTag(context *gin.Context) {
	// 获取id
	id := com.StrTo(context.Param("id")).MustInt()

	// 初始化验证
	valid := validation.Validation{}

	// 校验
	valid.Required(id, "id").Message("id 必须存在")
	valid.Min(id, 1, "id").Message("id 必须大于1")

	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagById(id) {
			models.DeleteTag(id)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": e.GetMsg(code),
		"data":    make(map[string]interface{}),
	})
}

func GetOneTag(context *gin.Context) {
	id := com.StrTo(context.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("id 必须大于1")

	var tag models.Tag
	var err error
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		tag, err = models.FindOneTag(id)
		if err != nil {
			code = e.ERROR_NOT_EXIST_TAG
			log.Println("require model error: ", err.Error())
			context.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": e.GetMsg(code),
			})
			return
		}
	} else {
		log.Println("has error : ", valid.Errors)
	}

	context.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": e.GetMsg(code),
		"data":    tag,
	})
}
