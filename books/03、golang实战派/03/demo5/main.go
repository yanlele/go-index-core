package main

import (
	"fmt"
	"log"
	"net/http"
)

/* 同时使用处理器和处理函数 */

func hiHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hi")
}

type WelcomeHandler struct {
	Name string
}

func (h WelcomeHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, h.Name)
}

func main() {
	mux := http.NewServeMux()

	// 注册处理函数
	mux.HandleFunc("/hi", hiHandler)

	mux.Handle("/welcome", WelcomeHandler{
		Name: "YanLe",
	})

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
