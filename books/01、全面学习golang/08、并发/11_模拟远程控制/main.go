package main

import (
	"errors"
	"fmt"
	"time"
)

// 模拟客户端接收信息
func RPCClient(ch chan string, req string) (string, error) {
	ch <- req
	select {
	case ack := <-ch:
		return ack, nil
	case <-time.After(time.Second):
		return "", errors.New("time out")
	}
}

func RPCServer(ch chan string) {
	for {
		data := <-ch
		fmt.Println("server received: ", data)
		ch <- "roger"
	}
}

func main() {
	ch := make(chan string)

	go RPCServer(ch)

	recv, err := RPCClient(ch, "hi")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("client received ", recv)
	}
}
