package app

import (
	//"controllers"
	"net/http"

	"github.com/sgarg2023/golangprojectpractice/mvc/controllers"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		panic(err)
	}
}
