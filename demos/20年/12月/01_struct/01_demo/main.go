package main

import "fmt"

type user struct {
	name string
	age  int
	address
}

type address struct {
	addressName string
	current     string
}

func main() {
	username := &user{}
	username.name = "yanle"
	username.age = 28
	userAddress := address{
		"四川",
		"成都",
	}
	username.address = userAddress

	fmt.Println(username)
	fmt.Println(username.address)
}
