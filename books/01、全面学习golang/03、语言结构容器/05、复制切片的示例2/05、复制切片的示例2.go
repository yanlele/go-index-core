package main

import "fmt"

func main() {
	const elementCount = 1000

	srcData := make([]int, elementCount)

	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}

	refData := srcData

	copyData := make([]int, elementCount)

	copy(copyData, srcData)
	srcData[0] = 999
	fmt.Println(refData[0])

	fmt.Println(copyData[0], copyData[elementCount - 1])

	// 复制原始数据从4到6(不包含)
	copy(copyData, srcData[4:6])
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", copyData[i])
	}
}
/*
999
0 999
4 5 2 3 4
*/
