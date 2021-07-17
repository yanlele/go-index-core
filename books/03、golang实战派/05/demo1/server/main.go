package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

/* 使用 json 编码实现的 PRC 服务端示例 */

func init() {
	fmt.Println("json 编码")
}

type ArgsLanguage struct {
	JAVA, GO string
}

type Programmer string

func (m *Programmer) GetSkill(all *ArgsLanguage, skill *string) error {
	*skill = "skill1: " + all.JAVA + ", skill2: " + all.GO
	return nil
}

func main() {
	str := new(Programmer)
	rpc.Register(str)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":8085")
	if err != nil {
		fmt.Println("ResolveTCPAddr err=", err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println("listener err = ", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		jsonrpc.ServeConn(conn)
	}
}
