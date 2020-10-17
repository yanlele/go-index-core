package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("goroutine 1")
	}()

	go func() {
		fmt.Println("goroutine 2")
	}()

	time.Sleep(time.Second)
}
