package controller

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	. "go-gorm-example/common/request"
	"go-gorm-example/common/response"
	"go-gorm-example/database"
	"go-gorm-example/models"
	"log"
	"net/http"
)

type AdminUser struct {
	response.CommonResponse
}

func (admin *AdminUser) QueryAllUser(context *gin.Context) {
	var users []models.AdminUser
	queryResult := database.DB.Select("id, name, username, created_at, updated_at").Order("id").Find(&users)
	if queryResult.Error != nil {
		admin.JsonFail(context, http.StatusBadRequest, "查询失败")
		return
	}
	admin.JsonSuccess(context, http.StatusOK, gin.H{"data": users})
}

func (admin *AdminUser) Store(context *gin.Context) {
	var request CreateRequest
	err := context.ShouldBind(&request)
	if err == nil {
		var count int64
		if result := database.DB.Model(&models.AdminUser{}).Where("username = ?", request.Username).Count(&count); result.Error != nil {
			admin.JsonFail(context, http.StatusBadRequest, result.Error.Error())
			return
		}
		if count > 0 {
			admin.JsonFail(context, http.StatusBadRequest, "用户名已经存在")
			return
		}

		// 密码加密
		password := []byte(request.Password)
		md5Context := md5.New()
		md5Context.Write(password)
		cipherString := md5Context.Sum(nil)
		user := models.AdminUser{
			Username: request.Username,
			Name:     request.Name,
			Password: hex.EncodeToString(cipherString),
		}

		// 保存失败
		if err := database.DB.Create(&user).Error; err != nil {
			admin.JsonFail(context, http.StatusBadRequest, err.Error())
			return
		}
		admin.JsonSuccess(context, http.StatusCreated, gin.H{"message": "创建成功"})
	} else {
		admin.JsonFail(context, http.StatusBadRequest, err.Error())
	}
}

func (admin *AdminUser) Update(context *gin.Context) {
	var request UpdateRequest
	err := context.ShouldBind(&request)
	if err == nil {
		var user models.AdminUser
		if database.DB.First(&user, context.Param("id")).Error != nil {
			admin.JsonFail(context, http.StatusNotFound, "数据不存在")
			return
		}
		user.Name = request.Name
		if err := database.DB.Save(&user).Error; err != nil {
			admin.JsonFail(context, http.StatusBadRequest, err.Error())
			return
		}
		admin.JsonSuccess(context, http.StatusOK, gin.H{})
		return
	}
	admin.JsonFail(context, http.StatusBadRequest, err.Error())
}

func (admin *AdminUser) Destroy(context *gin.Context) {
	var user models.AdminUser
	// 查询用户失败
	if database.DB.First(&user, context.Param("id")).Error != nil {
		admin.JsonFail(context, http.StatusNotFound, "数据不存在")
		return
	}

	// 删除失败
	if err := database.DB.Unscoped().Delete(&user).Error; err != nil {
		admin.JsonFail(context, http.StatusBadRequest, err.Error())
		return
	}

	admin.JsonSuccess(context, http.StatusCreated, gin.H{})
}

func (admin *AdminUser) FindOneUser(context *gin.Context) {
	var user models.AdminUser
	if database.DB.Select("id, name, username, created_at, updated_at").First(&user, context.Param("id")).Error != nil {
		admin.JsonFail(context, http.StatusNotFound, "数据不存在")
		return
	}
	admin.JsonSuccess(context, http.StatusOK, gin.H{"data": user})
}

/* 只是做一个get参数测试 */
func (admin *AdminUser) QueryTest(context *gin.Context) {
	query := context.Request.URL.Query()
	log.Println("query: ", query)

	name := query.Get("name")
	age := query.Get("age")
	address := query.Get("address")
	if address == "" {
		admin.JsonFail(context, http.StatusBadRequest, "address 为空")
		return
	}

	admin.JsonSuccess(context, http.StatusOK, gin.H{
		"query":   context.Request.URL.Query(),
		"name":    name,
		"age":     age,
		"address": address,
	})
}
