package main

import (
	"fmt"
	"go-gorm-example/config"
	_ "go-gorm-example/database"
	"go-gorm-example/routers"
)

func main() {
	conf := config.Get()
	// 答应地址
	fmt.Println("conf.DSN", conf.DSN)

	router := routers.InitRouter()
	_ = router.Run(config.Get().Addr)
}
