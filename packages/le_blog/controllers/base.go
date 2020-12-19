package controllers

import "github.com/gin-gonic/gin"

type Auth struct {
	Id         int
	Name       string
	Avatar     string
	Profession string
}

func (a Auth) GetAuth(c *gin.Context) Auth {
	auth, exists := c.Get("auth")
	if !exists {
		auth = Auth{
			Id:         0,
			Name:       "",
			Avatar:     "/static/logoh.png",
			Profession: "",
		}
	}
	return auth.(Auth)
}

type Header struct {
	Title string
}
