// finds the handler from the route and tries to execute the ServeHTTP function. Let us create
// a program called customMux.go and see this implementation in action
package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

type CustomServeMux struct {
}

func (p *CustomServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/" {
		giveRandom(w, r)
		return

	}
	http.NotFound(w, r)
	return
}

func giveRandom(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Your random number is: %f", rand.Float64())
}

func main() {

	mux := &CustomServeMux{}
	http.ListenAndServe(":8000", mux)
}
