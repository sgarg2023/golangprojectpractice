package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {

	alex := person{"Alex", "Anderson"}
	shubh := person{firstName: "Shubhra", lastName: "Garg"}
	fmt.Println(alex)
	fmt.Println(shubh)

}
