package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first cad to be 'Ace of spades' but got %v", d[0])
	}

	if d[len(d) - 1] != "Four of Clubs" {
		t.Errorf("Expected last card to be 'Four of clubs' but got %v", d[len(d) - 1])
	}
}

func TestSaveToFileAndDeckFromFile(t *testing.T) {

	os.Remove("_deckTesting")

	d := newDeck()
	d.saveToFile("_deckTesting")

	loadedDeck := deckFromFile("_deckTesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards but got %v", len(loadedDeck))
	}
	
	os.Remove("_deckTesting")
}
