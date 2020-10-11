package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{6, 7, 8}

	//copy(slice1, slice2)
	//fmt.Println(slice1)
	// [6 7 8 4 5]

	copy(slice2, slice1)
	fmt.Println(slice2)
	// [1 2 3]
}
