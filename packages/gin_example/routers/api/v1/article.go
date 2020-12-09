package v1

import (
	"gin-example/models"
	"gin-example/pkg/e"
	"gin-example/pkg/logging"
	"gin-example/pkg/setting"
	"gin-example/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

/* 获取单个文章 */
func GetArticle(context *gin.Context) {
	id := com.StrTo(context.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("id 必须大于 0")

	code := e.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			data = models.GetArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info("err.key: %s, err.message: %s\n", err.Key, err.Message)
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": e.GetMsg(code),
		"data":    data,
	})
}

func GetArticles(context *gin.Context) {
	// 初始化返回条件
	data := make(map[string]interface{})

	// 初始化查询条件
	maps := make(map[string]interface{})

	// 初始化验证
	valid := validation.Validation{}

	var state = -1
	// 获取 state
	if query := context.Query("state"); query != "" {
		state = com.StrTo(query).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
		maps["state"] = state
	}
	// 获取 tagid
	var tagId = -1
	if query := context.Query("tag_id"); query != "" {
		tagId = com.StrTo(query).MustInt()
		valid.Min(query, 1, "query").Message("tag id 必须大于1")
		maps["tag_id"] = tagId
	}

	// 如果不存在错误， 就分页查询
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		data["lists"], _ = models.GetArticles(util.GetPage(context), setting.AppSetting.PageSize, maps)
		data["total"], _ = models.GetArticleTotal(maps)
	} else {
		for _, err := range valid.Errors {
			logging.Info("err.key : %s, err.message : %s\n", err.Key, err.Message)
		}
	}

	// 返回json
	context.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": e.GetMsg(code),
		"data":    data,
	})
}

/*
新增文章
todo post 接口， 之后可以并 shouldBind 来改造
*/
func AddArticle(context *gin.Context) {
	tagId := com.StrTo(context.Query("tag_id")).MustInt()
	title := context.Query("title")
	desc := context.Query("desc")
	content := context.Query("content")
	createdBy := context.Query("created_by")
	state := com.StrTo(context.DefaultQuery("state", "0")).MustInt()

	// 封面图片， 可以非必填
	coverImageUrl := context.Query("cover_image_url")

	valid := validation.Validation{}
	valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(desc, "desc").Message("简述不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		if models.ExistTagById(tagId) {
			data := make(map[string]interface{})
			data["tag_id"] = tagId
			data["title"] = title
			data["desc"] = desc
			data["content"] = content
			data["created_by"] = createdBy
			data["state"] = state
			data["cover_image_url"] = coverImageUrl

			models.AddArticle(data)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info("err.key : %s, err.message : %s\n", err.Key, err.Message)
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": e.GetMsg(code),
		"data":    make(map[string]interface{}),
	})
}

/*
编辑文章

todo 用 shouldBind 来改造
*/
func EditArticle(context *gin.Context) {
	valid := validation.Validation{}

	id := com.StrTo(context.Param("id")).MustInt()
	tagId := com.StrTo(context.Query("tag_id")).MustInt()
	title := context.Query("title")
	desc := context.Query("desc")
	content := context.Query("content")
	modifiedBy := context.Query("modified_by")

	// 封面图片， 可以非必填
	coverImageUrl := context.Query("cover_image_url")

	var state int = -1
	if arg := context.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
	valid.MaxSize(desc, 255, "desc").Message("简述最长为255字符")
	valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")

	code := e.INVALID_PARAMS
	if valid.HasErrors() {
		var message string
		for _, err := range valid.Errors[0:1] {
			logging.Info("err.key: %s,  err.message: %s", err.Key, err.Message)
			message = err.Message
		}
		context.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": message,
		})
		return
	}

	if models.ExistArticleByID(id) && models.ExistTagById(tagId) {
		data := make(map[string]interface{})
		if tagId > 0 {
			data["tag_id"] = tagId
		}

		if title != "" {
			data["title"] = title
		}

		if desc != "" {
			data["desc"] = desc
		}

		if content != "" {
			data["content"] = content
		}

		if coverImageUrl != "" {
			data["cover_image_url"] = coverImageUrl
		}

		data["modified_by"] = modifiedBy
		models.EditArticle(id, data)
		code = e.SUCCESS
	}
	context.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": e.GetMsg(code),
		"data":    make(map[string]interface{}),
	})
}

func DeleteArticle(context *gin.Context) {
	id := com.StrTo(context.Param("id")).MustInt()
	valid := validation.Validation{}

	valid.Min(id, 1, "id").Message("id 必须大于0")

	code := e.INVALID_PARAMS
	if valid.HasErrors() {
		var message string
		for _, err := range valid.Errors[0:1] {
			message = err.Message
		}
		context.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": message,
		})
		return
	}

	if models.ExistArticleByID(id) {
		models.DeleteArticle(id)
		code = e.SUCCESS
	} else {
		code = e.ERROR_NOT_EXIST_ARTICLE
	}

	context.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": e.GetMsg(code),
		"data":    make(map[string]interface{}),
	})
}
