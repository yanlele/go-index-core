package main

import (
	"fmt"
	"github.com/yanlele/go-index-core/demos/21年/05月/02、获取当前是周几/date2week"
	"time"
)

func main() {
	// 获取当天时间是星期几
	fmt.Println(date2week.Date2Week())

	// 获取当天时间 + 24 小时是星期几
	fmt.Println(date2week.Date2Week(time.Now().Add(time.Hour * 24)))
}
