package main

import (
	"fmt"
	"net/rpc/jsonrpc"
)

type Send struct {
	JAVA, GO string
}

func main() {
	fmt.Println("client start ....")

	client, err := jsonrpc.Dial("tcp", "127.0.0.1:8085")
	if err != nil {
		fmt.Println("client dial error", err)
	}

	send := Send{"java", "go"}
	var receive string
	err = client.Call("Programmer.GetSkill", send, &receive)
	if err != nil {
		fmt.Println("call err", err)
	}

	fmt.Println("receive", receive)
}
