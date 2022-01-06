package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "ace of Spades" {
		t.Errorf("Expected ace of Spades, got %v", d[0])
	}
}

func TestSaveToFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := readFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		os.Remove("_decktesting")
		t.Errorf("length not right %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
