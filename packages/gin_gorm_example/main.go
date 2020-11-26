package main

import (
	"fmt"
	"go-gorm-example/config"
)

func main() {
	// 加载配置
	err := config.Load("config/config.yaml")
	if err != nil {
		fmt.Println("Failed to load configuration")
		return
	}

}
