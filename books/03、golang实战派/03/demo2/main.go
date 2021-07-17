package main

import (
	"fmt"
	"net/http"
)

/* 默认多路复用器指定多处理器 */
type handle1 struct {

}

// 名字必须要是 ServeHTTP 因为需要继承
/*
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
*/
func (h1 *handle1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi, handle1")
}

type handle2 struct {
}

func (h2 *handle2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi, handle2")
}

func main() {
	handle1 := handle1{}
	handle2 := handle2{}

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}

	http.Handle("/handle1", &handle1)
	http.Handle("/handle2", &handle2)
	server.ListenAndServe()
}
