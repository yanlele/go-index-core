package main

import "net/http"

/*
创建跳转 web 服务
*/

type Refer struct {
	handler http.Handler
	refer   string
}

func (refer *Refer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Referer() == refer.refer {
		refer.handler.ServeHTTP(w, r)
	} else {
		w.WriteHeader(403)
	}
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("this is handler"))
}

func hello(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("hello"))
}

func main() {
	referer := &Refer{
		handler: http.HandlerFunc(myHandler),
		refer:   "www.shirdon.com",
	}

	http.HandleFunc("/hello", hello)
	_ = http.ListenAndServe(":8080", referer)
}
