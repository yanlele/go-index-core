package main

import (
	"fmt"
	"sync"
	"time"
)

/* 用互斥锁的方式来试一试 demo3 看看要花费多少时间 */

var num int64
var wg sync.WaitGroup
var lock sync.Mutex

func write() {
	// 添加互斥锁
	lock.Lock()

	num = num + 1

	time.Sleep(10 * time.Millisecond)

	lock.Unlock()

	wg.Done()
}

func read() {
	lock.Lock()
	time.Sleep(time.Millisecond)
	lock.Unlock()
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()
	end := time.Now()

	fmt.Println("耗时： ", end.Sub(start))
}

/*耗时：  1.362264404s*/
