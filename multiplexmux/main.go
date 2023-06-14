package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func main() {

	newMux := http.NewServeMux()
	http.ListenAndServe(":8000", newMux)

	newMux.HandleFunc("/randomFloat", FloatHandler)
	newMux.HandleFunc("/randomInt", IntHandler)
}

func FloatHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, rand.Float64())
}

func IntHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, rand.Intn(100))
}
