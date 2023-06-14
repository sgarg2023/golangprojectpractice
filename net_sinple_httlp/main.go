package main

import (
	"io"
	"log"
	"net/http"
)

func MyServer(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Mike John!\n")
}

func main() {

	http.HandleFunc("/hello", MyServer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
