package main

import "fmt"

func main() {
	//var scene = make(map[string]int)
	scene := map[string]int{}

	scene["age"] = 28
	scene["name"] = 110

	delete(scene, "name")

	for key, value := range scene {
		fmt.Println(key)
		fmt.Println(value)
	}

	scene = make(map[string]int)

	for key, value := range scene {
		fmt.Println(key)
		fmt.Println(value)
	}
}
