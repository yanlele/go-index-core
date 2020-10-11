package main

import "fmt"

func main() {
	arr := []int{'y', 'a', 'n', 'l', 'e'}
	for index, value := range arr {
		fmt.Printf("index: %d, value: %d \n", index, value)
	}
}
