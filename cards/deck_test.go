package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length to be 2000 but received %v", len(d))
	}

	if d[0] != "Ace of Club" {
		t.Errorf("The first card expected is Ace of club but found %v", d[0])
	}

	if d[len(d)-1] != "Four of Heart" {
		t.Errorf("The last card oif the deck expected to be Four of Heart but found %v", d[len(d)-1])
	}
}

func TestFileToSaveAndNewDeckFromFile(t *testing.T) {

	os.Remove("__decktesting")
	newD := newDeck()
	newD.fileToSave("__decktesting")
	tryDeck := newDeckFromFile("__decktesting")

	if len(tryDeck) != 16 {
		t.Errorf("Expected length of deck to be 16 but received %v", len(tryDeck))
	}
	os.Remove("__decktesting")
}
