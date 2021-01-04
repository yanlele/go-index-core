package main

import (
	"fmt"
	"time"
)

func doneSelect(ch chan int) {
	for {
		fmt.Println("for 循环 - 输入 - start")
		select {
		case data := <-ch:
			fmt.Println(data)
			break
		default:
			fmt.Println("select default testing")
		}
		fmt.Println("for 循环 - 跳出 - end")
	}
}

func do() {
	ch := make(chan int)
	go doneSelect(ch)
	ch <- 10
}

func main() {
	do()
	time.Sleep(3 * time.Second)
}
