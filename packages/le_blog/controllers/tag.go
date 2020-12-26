package controllers

import (
	"github.com/gin-gonic/gin"
	"le-blog/bootstrap/driver"
	"le-blog/modules"
	"le-blog/utils"
	"net/http"
)

func TagIndex(c *gin.Context) {
	var tags []modules.Tag

	err := driver.DB.Find(&tags).Error

	if err != nil {
		utils.RedirectBack(c)
	}

	auth := Auth{}.GetAuth(c)

	data := &struct {
		Auth
		Tags []modules.Tag
	}{auth, tags}

	c.HTML(http.StatusOK, "tagIndex", data)
}

// 标签
func GetArticleByTagName(c *gin.Context) {
	tagName := c.Param("id")
	if tagName == "" {
		utils.RedirectBack(c)
		return
	}

	var articles []modules.Article
	//dbQuery := driver.Db.Table("articles").Where("find_in_set('" + tagName + "', tags)").Order("view_num")
	dbQuery := driver.DB.Table("articles").Where("find_in_set('" + tagName + "', tags)").Order("view_num")

	paginate, err := (&modules.Pagination{}).Paginate(c, *dbQuery, &articles)

	if err != nil {
		panic(err)
	}

	auth := Auth{}.GetAuth(c)
	header := Header{"标签"}
	data := struct {
		Paginate modules.Pagination
		Auth
		Header
	}{*paginate, auth, header}
	c.HTML(http.StatusOK, "index", data)
}

// GetTags 获取tags列表
func AjaxTags(c *gin.Context) {
	var tags []modules.Tag

	err := driver.DB.Select("id, name").Find(&tags).Error
	if err != nil {
		response := utils.Response{
			Status: 0,
			Data:   tags,
			Msg:    err.Error(),
		}
		c.JSON(http.StatusOK, response.FailResponse())
		return
	}
	c.JSON(http.StatusOK, tags)
	return
}

func AddTags(c *gin.Context) {
	var tag modules.Tag
	err := c.ShouldBind(&tag)
	if tag.Name == "" {
		res := utils.Response{
			Status: 401,
			Msg:    "标签名字一定要有",
		}
		c.JSON(http.StatusOK, res.FailResponse())
		return
	}

	// 绑定数据成功
	if err != nil {
		res := utils.Response{
			Status: 401,
			Msg:    err.Error(),
		}
		c.JSON(http.StatusOK, res.FailResponse())
		return
	}

	dbResult := driver.DB.Create(&tag)
	if dbResult.Error != nil {
		c.JSON(http.StatusOK, (&utils.Response{
			Msg: dbResult.Error.Error(),
		}).FailResponse())
		return
	}

	c.JSON(http.StatusOK, (&utils.Response{
		Msg: "保存成功",
	}).SuccessResponse())
}
