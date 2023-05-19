package main

import (
	//"encoding/json"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	FullName string `json:"full_name"`
	City     string `json:"city"`
	ZipCode  string `json:"zip_code"`
}

func myHandler(w http.ResponseWriter, c *http.Request) {
	fmt.Fprint(w, "Hello World, Welcome to the world of GOlang REST API")
}

func customers(w http.ResponseWriter, c *http.Request) {
	curt := []Customer{
		{FullName: "Rob", City: "Canada", ZipCode: "100094"},
		{FullName: "Bob", City: "Japan", ZipCode: "100095"},
		{FullName: "Micheal", City: "Australia", ZipCode: "100096"},
	}
	if c.Header.Get("Context-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(curt)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(curt)
	}
}
