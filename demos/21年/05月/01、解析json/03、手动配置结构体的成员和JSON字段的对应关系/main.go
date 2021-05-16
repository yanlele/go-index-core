package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Name string `json:"msg_name"`       // 对应JSON的msg_name
	Body string `json:"body,omitempty"` // 如果为空置则忽略字段
	Time int64  `json:"-"`              // 直接忽略字段
}

func main() {
	var m = Message{
		Name: "Alice",
		Body: "",
		Time: 1294706395881547000,
	}

	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println(string(data))

}
