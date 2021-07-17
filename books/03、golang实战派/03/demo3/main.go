package main

import (
	"fmt"
	"log"
	"net/http"
)

/* 自定义多路复用器 */

func hi(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hi web")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hi)

	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
