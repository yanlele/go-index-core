package controller

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	. "go-gorm-example/common/request"
	"go-gorm-example/common/response"
	"go-gorm-example/database"
	"go-gorm-example/models"
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
		admin.JsonSuccess(context, http.StatusOK, gin.H{"message": "创建成功"})
	} else {
		admin.JsonFail(context, http.StatusBadRequest, err.Error())
	}
}
