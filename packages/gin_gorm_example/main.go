package main

import (
	"fmt"
	"go-gorm-example/config"
	_ "go-gorm-example/database"
)

func main() {
	conf := config.Get()
	// 答应地址
	fmt.Println("conf.DSN", conf.DSN)

}
