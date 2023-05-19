package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {

	jim := person{
		firstName: "Jim",
		lastName:  "Corbert",
		contact: contactInfo{
			email:   "jim.corbert@gmail.com",
			zipCode: 56100,
		},
	}

	jim.print()
	fmt.Printf("\n")
	//jimPointer := &jim
	//jimPointer.updateStruct("Fluke")
	jim.updateStruct("Sandra")
	fmt.Printf("\n")
	fmt.Printf("\n Willing to see whether Jim's firstName got updated or not\n")
	jim.print()

	mySlice := []string{"Hi", "I", "am", "shubhra", "garg"}
	fmt.Printf("\n-------------------Before slice updation------------------------\n")
	fmt.Println(mySlice)
	updateSlice(mySlice)

	fmt.Printf("\n-------------------After slice updation------------------------\n")
	fmt.Println(mySlice)

}

func updateSlice(s []string) {
	s[0] = "Bye"
}

func (personPointer *person) updateStruct(newFirstName string) {
	(*personPointer).firstName = newFirstName
	//fmt.Printf("\n Willing to see the receiver copy value\n")
	//p.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
