package main

import "fmt"

func playerGen(name string) func() (string, int) {
	hp := 50
	return func() (string, int) {
		return name, hp
	}
}

func main() {
	generator := playerGen("yanle")
	name, ph := generator()
	fmt.Println(name, ph)
}
