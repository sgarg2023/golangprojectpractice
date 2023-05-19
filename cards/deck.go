package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type deck []string

func (d deck) printDeck() {

	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Club", "Diamonds", "Spades", "Heart"}
	cardValue := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValue {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) fileToSave(filename string) {
	os.WriteFile(filename, []byte(d.toString()), 0744)
}

func newDeckFromFile(filename string) deck {

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("ERROR :", err)
		os.Exit(1)
	}

	s := strings.Split(string(data), ",")
	return deck(s)

}

func (d deck) shuffle() {

	for i := range d {
		newPos := rand.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}

}
