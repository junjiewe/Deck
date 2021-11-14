package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 27 {
		t.Errorf("Expected deck length of 27, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spade, but got %v", d[0])
	}

	if d[len(d)-1] != "Nine of Diamonds" {
		t.Errorf("Expected last card of Nine of Diamonds, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.savetoFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 27 {
		t.Errorf("Expected 27 cards in deck, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
