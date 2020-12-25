package main

import "fmt"

type User struct {
	name string
	age  int
}

func (user *User) get() {
	fmt.Println("name: ", user.name)
	fmt.Println("age: ", user.age)
}

func (user User) getUser() {
	fmt.Println("user name :", user.name)
}

func main() {
	user := User{
		"YANLE",
		29,
	}

	(&User{
		"Yanlele",
		29,
	}).get()

	user.get()

	User{
		age: 29,
	}.getUser()
}
