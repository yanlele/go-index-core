package main

import (
	"log"
	"net/http"
)

/* 创建一个简单的 https 服务器 */

func main() {
	ser := &http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(handle),
	}
	log.Printf("serving on http://127.0.0.1:8080")
	log.Fatal(ser.ListenAndServe())
}

func handle(w http.ResponseWriter, r *http.Request) {
	log.Printf("got connection %s", r.Proto)
	w.Write([]byte("hello this is a http message!"))
}
