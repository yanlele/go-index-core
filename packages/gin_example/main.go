package main

import (
	"fmt"
	"gin-example/models"
	"gin-example/pkg/gredis"
	"gin-example/pkg/logging"
	"gin-example/pkg/setting"

	"gin-example/routers"
	"github.com/fvbock/endless"
	"log"
	"syscall"
)

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	_ = gredis.SetUp()
}

func main() {
	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}

	//router := routers.InitRouter()
	//_ = router.Run(fmt.Sprintf(":%d", setting.HTTPPort))

	//s := &http.Server{
	//	Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
	//	Handler:        router,
	//	ReadTimeout:    setting.ReadTimeout,
	//	WriteTimeout:   setting.WriteTimeout,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//
	//s.ListenAndServe()
}
