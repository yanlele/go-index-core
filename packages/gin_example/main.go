package main

import (
	"fmt"
	"gin-example/pkg/setting"
	"gin-example/routers"
)

func main() {
	router := routers.InitRouter()

	_ = router.Run(fmt.Sprintf(":%d", setting.HTTPPort))
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
