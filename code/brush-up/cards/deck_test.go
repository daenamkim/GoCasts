package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	// t is a test handler
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	name := "_deck_save_test"

	os.Remove(name)

	d := newDeck()
	d.saveToFile(name)
	loadedDeck := newDeckFromFile(name)

	if len(loadedDeck) != len(d) {
		t.Errorf("Expected %v cards in deck, but got %v", len(d), len(loadedDeck))
	}

	os.Remove(name)
}
