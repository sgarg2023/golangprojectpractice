package main

import (
	"fmt"

	"github.com/sgarg2023/golangprojectpractice/toolkit"
)

func main() {
	var tools toolkit.Tools
	s := tools.randomString(10)
	fmt.Println("Random String:", s)
}
