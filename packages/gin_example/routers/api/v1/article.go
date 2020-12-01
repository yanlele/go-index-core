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
			log.Printf("err.key: %s, err.message: %s\n", err.Key, err.Message)
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
		data["lists"], _ = models.GetArticles(util.GetPage(context), setting.PageSize, maps)
		data["total"], _ = models.GetArticleTotal(maps)
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key : %s, err.message : %s\n", err.Key, err.Message)
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

			models.AddArticle(data)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key : %s, err.message : %s\n", err.Key, err.Message)
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": e.GetMsg(code),
		"data":    make(map[string]interface{}),
	})
}

func EditArticle(context *gin.Context) {

}

func DeleteArticle(context *gin.Context) {

}
