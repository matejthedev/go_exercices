package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 20 {
		t.Errorf("Expected deck length of 20, got %v.", len(d))
	}

	if d[0] != "Ace of spades" {
		t.Errorf("Expected Ace of spades, got %v.", d[0])
	}

	if d[len(d)-1] != "Five of hearts" {
		t.Errorf("Expected Five of hearts, got %v.", d[len(d)-1])
	}
}

func TestSaveToFileAndGetNewDeckFromFile(t *testing.T) {
	os.Remove("_decktest")
	d := newDeck()
	d.saveToFile("_decktest")
	deck := getDeckFromFile("_decktest")
	if len(deck) != 20 {
		t.Errorf("Ecpected length of 20, got %v", len(deck))
	}
	os.Remove("_decktest")
}
