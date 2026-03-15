package main

import (
	"net/http"
)

func mainHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("It's main Page"))
}

func apiHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("It's api page"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", mainHandler)
	mux.HandleFunc("/api/", apiHandler)

	err := http.ListenAndServe("localhost:3300", mux)
	if err != nil {
		panic(err)
	}
}
