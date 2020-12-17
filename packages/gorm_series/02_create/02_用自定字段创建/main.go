package main

import (
	"fmt"
	"go-index/packages/gorm_series/initDB"
	"go-index/packages/gorm_series/model"
	"time"
)

func main() {
	student := model.Student{
		Name:     "jinzhu",
		Age:      18,
		Birthday: time.Now().UnixNano() / 1e6,
	}

	// 创建记录并更新给出的字段。
	result := initDB.DB.Select("Name", "Age", "Birthday").Create(&student)

	if result.Error != nil {
		fmt.Println("创建失败", result.Error.Error())
	}
	// 插入的主键
	fmt.Println("student.id", student.Id)
	// 返回插入记录的条数
	fmt.Println("更新条数： ", result.RowsAffected)

	// 创建记录并更新未给出的字段。
	result = initDB.DB.Omit("Name", "Age", "Birthday").Create(&student)

	if result.Error != nil {
		fmt.Println("创建失败", result.Error.Error())
	}
	// 插入的主键
	fmt.Println("student.id", student.Id)
	// 返回插入记录的条数
	fmt.Println("更新条数： ", result.RowsAffected)
}
