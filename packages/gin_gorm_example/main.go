package main

import (
	"go-gorm-example/config"
	_ "go-gorm-example/database"
	"go-gorm-example/routers"
)

func main() {
	router := routers.InitRouter()
	_ = router.Run(config.Get().Addr)
}
