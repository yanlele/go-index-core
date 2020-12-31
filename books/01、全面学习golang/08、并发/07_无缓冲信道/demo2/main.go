package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func handleRunner(baton chan int) {
	var newRunner int
	runner := <-baton
	fmt.Printf("Runner %d Running With Baton\n", runner)
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To The Line\n", newRunner)
		go handleRunner(baton)
	}

}

func main() {

}
