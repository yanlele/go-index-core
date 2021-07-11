package main

import (
	"fmt"
	"sync"
)

/*生活中最明显的例子就是，大家抢着上厕所，资源有限，只能一个一个的用*/

var num int64
var wg sync.WaitGroup

func add() {
	for i := 0; i < 10000000; i++ {
		num = num + 1
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
