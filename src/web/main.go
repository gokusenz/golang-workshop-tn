package main

import (
	"fmt"
	"net/http"
)

func response(rw http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	rw.Write([]byte("Hello world."))
}

func fail(rw http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	rw.WriteHeader(500)
	rw.Write([]byte("Fail"))
}

func main() {
	http.HandleFunc("/", response)
	http.HandleFunc("/fail", fail)
	http.ListenAndServe(":8080", nil)
}
