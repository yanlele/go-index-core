package main

import (
	"fmt"
	"sort"
)

func main() {
	scene := make(map[string]int)
	var sceneList []string

	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960

	for key, value := range scene {
		fmt.Printf("key: %s, value: %d\n", key, value)
		sceneList = append(sceneList, key)
	}

	sort.Strings(sceneList)
	fmt.Println(sceneList)
}

/*
key: route, value: 66
key: brazil, value: 4
key: china, value: 960
[brazil china route]
*/
