package main

import (
	"fmt"
	"net/http"
)

func response(rw http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	rw.Write([]byte("Hello world."))
}

func main() {
	http.HandleFunc("/", response)
	http.ListenAndServe(":8080", nil)
}
