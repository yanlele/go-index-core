package main

import (
	"fmt"
	"sync"
)

/*互斥锁*/

var num int64
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	for i := 0; i < 10000000; i++ {
		// 访问前给资源加锁
		lock.Lock()
		num = num + 1
		lock.Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(2)

	go add()
	go add()

	wg.Wait()
	fmt.Println(num)
}
