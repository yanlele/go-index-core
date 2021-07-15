package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
创建一个 get 请求

获取百度 html 文档
*/

func main() {
	res, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("err", err)
	}
	closer := res.Body
	bytes, err := ioutil.ReadAll(closer)
	fmt.Println(string(bytes))
}
