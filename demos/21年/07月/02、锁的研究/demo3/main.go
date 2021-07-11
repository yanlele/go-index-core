package main

import (
	"fmt"
	"sync"
	"time"
)

/*读写锁*/

var (
	num    int64
	wg     sync.WaitGroup
	rwLock sync.RWMutex
)

func write() {
	// 加上写锁
	rwLock.Lock()

	num = num + 1

	time.Sleep(10 * time.Millisecond)

	rwLock.Unlock()

	wg.Done()
}

func read() {
	rwLock.RLock()

	time.Sleep(time.Millisecond)

	rwLock.RUnlock()

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
	end:=time.Now()
	fmt.Println(end.Sub(start))
}
