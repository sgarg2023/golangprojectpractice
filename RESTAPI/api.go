package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		handleGreet(w, r)
	})
	fmt.Printf("Server started on address %s", s.addr)
	http.ListenAndServe(s.addr, router)
}
