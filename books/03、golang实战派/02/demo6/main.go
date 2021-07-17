package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

/* 服务端创建 PUT 请求 */

func main() {
	url := ""
	payload := strings.NewReader("")

	req, _ := http.NewRequest("DELETE", url, payload)
	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
