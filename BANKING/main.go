package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	name    string `json: "name"`
	City    string `json: "City"`
	zipCode string `json: "zipCode"`
}

func main() {

	http.HandleFunc("/greet", myHandler)
	http.HandleFunc("/customers", customers)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func myHandler(w http.ResponseWriter, c *http.Request) {
	fmt.Fprint(w, "Hello Shubhra, Welcome to the world of GOlang REST API")
}

func customers(w http.ResponseWriter, c *http.Request) {
	cust := []Customer{
		//{name: "Alice", City: "Canada", zipCode: "100075"},
		//{name: "Bob", City: "Japan", zipCode: "100074"},
		//{name: "Marshall", City: "Australia", zipCode: "100076"},
		{"Alice", "Canada", "100075"},
		{"Bob", "Australia", "100076"},
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cust)
}
