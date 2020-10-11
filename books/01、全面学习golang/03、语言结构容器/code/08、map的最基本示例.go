package main

import "fmt"

func main() {
	// 定义方式1
	var mapAssigned map[string]int
	mapLit := map[string]int{"one": 1, "tow": 2}
	// 定义方式2
	mapCreated := make(map[string]float32)
	mapAssigned = mapLit
	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapAssigned["two"] = 3

	fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapLit["two"])

	// 如果没有的话， 会后默认值
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])
}

/*
Map literal at "one" is: 1
Map created at "key2" is: 3.141590
Map assigned at "two" is: 3
Map literal at "ten" is: 0
*/
