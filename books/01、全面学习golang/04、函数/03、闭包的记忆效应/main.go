package main

import "fmt"

// 典型闭包应用而已
func accumulate(value int) func() int {
	return func() int {
		value++
		return value
	}
}

func main() {
	accumulator := accumulate(1)
	fmt.Println(accumulator())
	fmt.Println(accumulator())

	// 打印累加器的函数地址
	fmt.Printf("%p\n", &accumulator)

	// 新创建一个累加器
	accumulator2 := accumulate(10)

	// 累加
	fmt.Println(accumulator2())

	// 打印累加器的函数地址
	fmt.Printf("%p\n", &accumulator2)
}
