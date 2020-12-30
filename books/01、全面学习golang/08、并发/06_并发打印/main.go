package main

import "fmt"

var printer = func(c chan int) {
	for {
		data := <-c
		if data == 0 {
			break
		}
		// 打印数据
		fmt.Println(data)
	}
	c <- 0
}

func main() {
	c := make(chan int)
	go printer(c)
	for i := 0; i <= 10; i++ {
		c <- i
	}

	c <- 0
	<-c
}
