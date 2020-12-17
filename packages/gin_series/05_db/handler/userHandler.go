package handler

import (
	"github.com/gin-gonic/gin"
	"go-index/packages/gin_series/05_db/model"
	"log"
	"net/http"
)

func UserSave(context *gin.Context) {
	username := context.Param("name")
	context.String(http.StatusOK, "用户"+username+"已经保存")
}

// 通过 query 方法进行获取参数
func UserSaveByQuery(context *gin.Context) {
	username := context.Query("name")
	//age := context.Query("age")
	age := context.DefaultQuery("age", "20")
	context.String(http.StatusOK, "用户:"+username+",年龄:"+age+"已经保存")
}

func UserRegister(context *gin.Context) {
	var user model.UserModel
	if err := context.ShouldBind(&user); err != nil {
		context.String(http.StatusBadRequest, "输入的数据不合法")
		log.Panicln("err ->", err.Error())
	}
	passwordAgain := context.PostForm("password-again")
	if passwordAgain != user.Password {
		context.String(http.StatusBadRequest, "密码校验无效，两次密码不一致")
		log.Panicln("密码校验无效，两次密码不一致")
	}
	id := user.Save()
	log.Println("id is ", id)
	context.Redirect(http.StatusMovedPermanently, "/")
}

func UserLogin(context *gin.Context) {
	var user model.UserModel
	if err := context.Bind(&user); err != nil {
		log.Panicln("login 绑定错误", err.Error())
	}

	currentUser := user.QueryByEmail()
	log.Println("currentUser: ", currentUser.Email)
	if currentUser.Password == user.Password {
		log.Println("登录成功", currentUser.Email)
		context.HTML(http.StatusOK, "index.tmpl", gin.H{
			"email": currentUser.Email,
		})
	}
}
