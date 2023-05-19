package main

import "fmt"

func main() {
	//use of slice
	card := newDeck()

	dealDeck, otherDeck := deal(card, 5)

	fmt.Println("-----------------Print deal deck----------------")
	dealDeck.printDeck()
	fmt.Println("-----------------Print other deck---------------")
	otherDeck.printDeck()
	card.printDeck()

	//testing toString() function with a new deck of cards
	myCard := newDeck()
	fmt.Println("I am shubhra garg--------------------------")
	//fmt.Println(myCard.toString())
	myCard.fileToSave("promisingcards")

	mydeck := newDeckFromFile("promisingcards")
	fmt.Println("The deck from the saved file is ----------------------------")
	mydeck.printDeck()

	fmt.Println("--------Lets try Shuffling of cards----------------")
	newD := newDeck()
	newD.shuffle()
	newD.printDeck()

}
