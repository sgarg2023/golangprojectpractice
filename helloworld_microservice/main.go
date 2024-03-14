package main

import (
	"fmt"
	"net/http"
)

func main() {

	port := 9030
	http.HandleFunc("/helloworld", HelloWorldHandler)
	fmt.Println("starting the server")
	//log.Printf("Server starting at port : %v", port)
	//log.fatal(http.ListenAndServe(fmt.Sprintf("%v",port),nil))
	http.ListenAndServe(fmt.Sprintf("%v", port), nil)
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world shubhra garg\n")
}
