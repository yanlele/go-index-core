package main

import (
	"fmt"
	"sync"
)

func main() {
	var scene sync.Map
	scene.Store("name", "yanle")
	scene.Store("age", 27)
	scene.Store("address", "重庆")

	name, _ := scene.Load("name")
	fmt.Println(name)

	scene.Delete("age")

	scene.Range(func(key, value interface{}) bool {
		fmt.Println("key: ", key)
		fmt.Println("value: ", value)
		return true
	})
}
