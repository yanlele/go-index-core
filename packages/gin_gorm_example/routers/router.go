package routers

import (
	"github.com/gin-gonic/gin"
	"go-gorm-example/controller"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		adminUser := &controller.AdminUser{}
		api.GET("/admin_users", adminUser.QueryAllUser)
		api.POST("/admin_user", adminUser.Store)
		api.PATCH("/admin_user/:id", adminUser.Update)
		api.DELETE("/admin_user/:id", adminUser.Destroy)
	}

	return router
}
