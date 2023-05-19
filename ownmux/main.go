package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	FullName   string
	City       string
	ZipCode    string
	CustomerId int
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/greet", MyHandler).Methods("GET")
	router.HandleFunc("/customers", Customers).Methods("GET")
	router.HandleFunc("/customers/{CustomerId}", GetCustomer).Methods("GET")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

func MyHandler(w http.ResponseWriter, c *http.Request) {
	fmt.Fprint(w, "Hello World, Welcome to the world of GOlang REST API")
}

func GetCustomer(w http.ResponseWriter, c *http.Request) {
	fmt.Fprint(w, "Get customer API is being called up, thanks!\n")
	vars := mux.Vars(c)
	fmt.Fprint(w, vars["CustomerId"])
}

func Customers(w http.ResponseWriter, c *http.Request) {
	curt := []Customer{
		{FullName: "Alice", City: "Canada", ZipCode: "100090", CustomerId: 12345},
		{FullName: "Bob", City: "Japan", ZipCode: "100095", CustomerId: 90000},
		{FullName: "Marshall", City: "Australia", ZipCode: "100096", CustomerId: 1234578},
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(curt)
}
