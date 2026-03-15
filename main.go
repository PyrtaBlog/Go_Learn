package main

import (
	"net/http"
)

type MyHandler struct {
}

func (h MyHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	data := []byte("Hello")
	res.Write(data)
}

func main() {
	var h MyHandler

	err := http.ListenAndServe("localhost:3300", h)
	if err != nil {
		panic(err)
	}
}
