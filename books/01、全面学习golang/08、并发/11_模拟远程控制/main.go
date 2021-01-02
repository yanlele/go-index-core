package main

import (
	"errors"
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

func main() {
}
