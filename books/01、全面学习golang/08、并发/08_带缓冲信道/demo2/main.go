package main

import "fmt"

func demo(count chan int) {
	for {
		count1 := <-count
		count2 := <-count

		if count1 == 2 {
			break
		}
		fmt.Println("count2", count2)
		fmt.Println("count1", count1)
	}
}

func main() {
	ch := make(chan int, 2)
	ch <- 1

	go demo(ch)

	ch <- 2

	<- ch
}
