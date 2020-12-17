package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"le-blog/bootstrap"
	"le-blog/bootstrap/driver"
	"le-blog/config"
	"log"
)

var app *gin.Engine

func init() {
	config.InitConfig()
	driver.InitConn()
}

func main() {
	env := config.Config.Section("env")
	port, err := env.GetKey("Port")
	if err != nil {
		panic(err)
	}

	gin.ForceConsoleColor()
	app = bootstrap.Init()

	log.Fatal(app.Run(fmt.Sprintf(":%s", port)))
}
