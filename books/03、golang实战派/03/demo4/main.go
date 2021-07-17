package main

import (
	"fmt"
	"log"
	"net/http"
)

/* 自定义多路复用器 - 多路由匹配 */

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "首页， indexHandler")
}

func hiHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hi, hiHandler")
}

func webHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "web, webHandler")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/hi", hiHandler)
	mux.HandleFunc("/web", webHandler)

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
