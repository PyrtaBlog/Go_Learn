package main

import (
	"net/http"
)

func main() {
	err := http.ListenAndServe("localhost:3300", nil)
	if err != nil {
		panic(err)
	}
}
