package main

import (
	"fmt"
	"sync"
)

/*
sync.WaitGroup

WaitGroup 内部实现了一个计数器，用来记录未完成的操作个数，它提供了三个方法：

- Add() 用来添加计数
- Done() 用来在操作结束时调用，使计数减一 【我不会告诉你 Done() 方法的实现其实就是调用 Add(-1)】
- Wait() 用来等待所有的操作结束，即计数变为 0，该函数会在计数不为 0 时等待，在计数为 0 时立即返回

*/
func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("goroutine 1")
		wg.Done()
	}()

	go func() {
		fmt.Println("goroutine 2")
		wg.Done()
	}()

	wg.Wait()
}
