package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Excepted deck length of 20, but got %v", len(d))
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDcckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Excepted deck length of 20, but got %v", len("loadedDeck"))
	}

	os.Remove("_decktesting")
}
