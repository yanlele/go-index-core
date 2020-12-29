package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter = 10

func countFunction(lock *sync.Mutex) {
	lock.Lock()
	counter++
	fmt.Println(counter)
	lock.Unlock()
}

func main() {
	lock := &sync.Mutex{}
	fmt.Println("temp1")
	for i := 0; i < 10; i++ {
		go countFunction(lock)
	}
	fmt.Println("temp2")

	for {
		lock.Lock()
		fmt.Println("temp3")
		c := counter
		lock.Unlock()
		runtime.Gosched()
		fmt.Println("main counter: ", c)
		if c >= 10 {
			break
		}
	}
}
