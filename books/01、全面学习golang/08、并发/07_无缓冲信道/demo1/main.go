package main

import (
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func player(name string, court chan int) {
	defer wg.Done()
}

func main() {

}
