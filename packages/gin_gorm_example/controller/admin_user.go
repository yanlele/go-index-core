package controller

import (
	"github.com/gin-gonic/gin"
	"go-gorm-example/controller/common"
	"go-gorm-example/database"
	"go-gorm-example/models"
	"net/http"
)

type AdminUser struct {
	common.Common
}

func (admin *AdminUser) QueryAllUser(context *gin.Context) {
	var users []models.AdminUser
	queryResult := database.DB.Select("id, name, username, created_at, updated_at").Order("id").Find(&users)
	if queryResult.Error != nil {
		admin.JsonFail(context, http.StatusOK, "查询失败")
		return
	}
	admin.JsonSuccess(context, http.StatusOK, gin.H{"data": users})
}
