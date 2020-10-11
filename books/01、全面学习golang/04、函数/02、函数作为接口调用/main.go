package main

import "fmt"

type Invoker interface {
	call(interface{})
}

type Type struct {
}

func (s *Type) call(p interface{}) {
	fmt.Println("form struct: ", p)
}

type FuncCaller func(interface{})

func (f FuncCaller) call(p interface{}) {
	f(p)
}

func main() {
	var invoker Invoker
	s := new(Type)

	invoker = s

	invoker.call("hello")
}
