package main

import (
	//"encoding/json"
	"log"
	"net/http"
)

func Start() {
	http.HandleFunc("/greet", myHandler)
	http.HandleFunc("/customers", customers)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
