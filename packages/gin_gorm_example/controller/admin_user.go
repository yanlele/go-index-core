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

func (admin *AdminUser) Index(context *gin.Context) {
	var users []models.AdminUser
	database.DB.Select("id, name, username, create_at, update_at").Order("id").Find(&users)
	admin.JsonSuccess(context, http.StatusOK, gin.H{"data": users})
}
