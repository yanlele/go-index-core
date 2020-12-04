package api

import (
	"gin-example/models"
	"gin-example/pkg/e"
	"gin-example/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(context *gin.Context) {
	username := context.Query("username")
	password := context.Query("password")

	valid := validation.Validation{}
	ok, _ := valid.Valid(&auth{username, password})

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		id, _ := models.CheckAuth(username, password)
		if id > 0 {
			token, err := util.GenerateToken(username, password, id)
			code = e.ERROR_AUTH_TOKEN

			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": e.GetMsg(code),
		"data":    data,
	})
}
