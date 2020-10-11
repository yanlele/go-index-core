package main

import "fmt"

func main() {
	var arr = []int{1, 2, 3, 4, 5}

	// 删除第一个元素
	fmt.Println(arr[1:])

	// 删除前三个元素
	fmt.Println(arr[3:])

	// 删除第三个元素
	fmt.Println(append(arr[:2], arr[3:]...))

	// 删除第三四个元素
	fmt.Println(append(arr[:2], arr[4:]...))

	// 删除最后一个元素
	fmt.Println(arr[:len(arr)-1])
	fmt.Println(arr[:len(arr)-2])
}
/*
[2 3 4 5]
[4 5]
[1 2 4 5]
[1 2 5]
[1 2 5 5]
[1 2 5]
*/
