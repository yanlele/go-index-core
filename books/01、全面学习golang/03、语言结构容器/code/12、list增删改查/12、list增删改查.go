package main

import (
	"container/list"
	"fmt"
)

func main() {
	myList := list.New()
	myList.PushBack("yanle")
	myList.PushFront("name")

	element := myList.PushBack("lele")
	myList.InsertAfter("insertAfter", element)
	myList.InsertBefore("insertBefore", element)

	myList.Remove(element)

	// 遍历
	for i := myList.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
