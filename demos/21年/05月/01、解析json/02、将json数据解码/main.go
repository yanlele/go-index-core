package main

import (
	"encoding/json"
	"fmt"
)

var jsonBlob = []byte(`[
{"Name": "Platypus", "Order": "Monotremata"},
{"Name": "Quoll", "Order": "Dasyuromorphia"}
]`)

type Animal struct {
	Name  string
	Order string
}

func main() {
	var animals []Animal

	// func Unmarshal(data []byte, v interface{}) error
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Printf("%+v", animals)
}
