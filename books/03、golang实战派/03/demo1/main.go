package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloWord(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello go web")
}

func main() {
	http.HandleFunc("/hello", helloWord)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
