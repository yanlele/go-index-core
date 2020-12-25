package controllers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"le-blog/bootstrap/driver"
	"le-blog/modules"
	"le-blog/utils"
	"net/http"
)

// 表单提交数据
type formUser struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
}

// 注册页面
func Register(c *gin.Context) {
	auth := Auth{}.GetAuth(c)
	if auth.Id > 0 {
		utils.Redirect(c, "/")
	}

	data := struct {
		Title string
		Auth
	}{"注册", auth}

	c.HTML(http.StatusOK, "join", data)
}

// 注册用户
func DoRegister(c *gin.Context) {
	var regData formUser
	if err := c.ShouldBind(&regData); err != nil {
		response := utils.Response{
			Status: 403,
			Data:   nil,
			Msg:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response.FailResponse())
		return
	}
	var user modules.User

	// 保存数据到数据库
	driver.DB.Where("email = ?", regData.Email).First(&user)

	if user.ID != 0 {
		res := utils.Response{
			Status: 1001,
			Data:   nil,
			Msg:    "邮箱已经存在",
		}
		c.JSON(http.StatusOK, res.FailResponse())
		return
	}

	salt := utils.Salt()
	user = modules.User{
		Name:       regData.Name,
		Password:   utils.CryptUserPassword(regData.Password, salt),
		Salt:       salt,
		Email:      regData.Email,
		Profession: "",
		Avatar:     "",
	}

	// 保存数据
	err := driver.DB.Create(&user).Error
	if err != nil {
		res := utils.Response{
			Status: 500,
			Data:   nil,
			Msg:    err.Error(),
		}
		c.JSON(http.StatusOK, res.FailResponse())
		return
	}

	// 将用户数据写入session
	auth := &Auth{
		Id:         int(user.ID),
		Name:       user.Name,
		Avatar:     user.Avatar,
		Profession: user.Profession,
	}

	session := sessions.Default(c)
	session.Set("auth", auth)
	err = session.Save()

	fmt.Println(session.Get("auth"))

	if err != nil {
		res := utils.Response{
			Status: 500,
			Data:   nil,
			Msg:    err.Error(),
		}
		c.JSON(http.StatusOK, res.FailResponse())
		return
	}

	// 成功的场景
	res := utils.Response{
		Data: regData,
	}
	c.JSON(http.StatusOK, res.SuccessResponse())
}

// get sign in
func Login(c *gin.Context) {
	auth := Auth{}.GetAuth(c)
	if auth.Id != 0 {
		utils.Redirect(c, "/")
	}

	data := struct {
		Auth
	}{auth}
	c.HTML(http.StatusOK, "login", data)
}

func DoLogin(c *gin.Context) {
	var logData struct {
		Email    string `form:"email" json:"email" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}

	// 验证数据
	if err := c.ShouldBind(&logData); err != nil {
		res := utils.Response{
			Status: 403,
			Msg:    err.Error(),
		}
		c.JSON(http.StatusOK, res.FailResponse())
		return
	}

	// 查找用户看密码是否匹配
	user := modules.User{}
	err := driver.DB.Where("email = ?", logData.Email).First(&user).Error
	if err != nil {
		c.JSON(http.StatusOK, (&utils.Response{
			Status: 403,
			Data:   nil,
			Msg:    "邮箱或者密码错误",
		}).FailResponse())
		return
	}

	// 查看是否匹配
	ok := utils.VerifyUserPassword(&user, logData.Password)
	if !ok {
		c.JSON(http.StatusOK, (&utils.Response{
			Status: 403,
			Data:   nil,
			Msg:    "邮箱或者密码错误",
		}).FailResponse())
		return
	}

	auth := &Auth{
		Id:         int(user.ID),
		Name:       user.Name,
		Avatar:     user.Avatar,
		Profession: user.Profession,
	}

	session := sessions.Default(c)
	session.Set("auth", auth)
	err = session.Save()
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, (&utils.Response{
		Status: 0,
		Data:   nil,
		Msg:    "",
	}).SuccessResponse())
	return
}
