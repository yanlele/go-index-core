package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4}
	println(a[0])
	println(a[len(a)-1])
	for index, value := range a {
		fmt.Printf("index: %d, value: %d \n", index, value)
	}
}

/*
1
4
index: 0, value: 1
index: 1, value: 2
index: 2, value: 3
index: 3, value: 4
*/
