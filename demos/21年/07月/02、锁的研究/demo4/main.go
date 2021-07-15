package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/* 原子操作的示例 */

var num int64
var lock sync.Mutex
var wg sync.WaitGroup

// 普通版本
func add() {
	num = num + 1
	wg.Done()
}

// 互斥锁版本
func mutexAdd() {
	lock.Lock()
	num = num + 1
	lock.Unlock()
	wg.Done()
}

func atomicAdd() {
	atomic.AddInt64(&num, 1)
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 20000; i++ {
		wg.Add(1)
		//go add()
		//go mutexAdd()
		go atomicAdd()
	}
	wg.Wait()

	end := time.Now()
	fmt.Println("num: ", num)
	fmt.Println(end.Sub(start))
}
