package main

import (
	"fmt"
	"go-index/packages/gorm_series/initDB"
	"go-index/packages/gorm_series/model"
	"log"
	"time"
)

func main() {
	student := model.Student{
		Name:     "jinzhu",
		Age:      18,
		Birthday: time.Now().UnixNano() / 1e6,
	}
	result := initDB.DB.Create(&student)
	if result.Error != nil {
		log.Panicln("创建失败", result.Error.Error())
	}
	// 插入的主键
	fmt.Println(student.Id)
	// 返回插入记录的条数
	fmt.Println(result.RowsAffected)
}
