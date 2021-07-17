package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

/* 服务端创建 post 请求 */

func main() {
	url := ""
	body := ""
	res, err := http.Post(url, "", bytes.NewBuffer([]byte(body)))

	if err != nil {
		fmt.Println("err", err)
	}

	buffer, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(buffer))
}
