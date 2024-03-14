package main

import (
	"fmt"
	"path"
)

func main() {

	var dir, file string
	dir, file = path.Split("D:/goprojects/httprouter/main.go")
	fmt.Println("dir:", dir)
	fmt.Println("filename:", file)

}
