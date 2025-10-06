package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %d", len(d))
	}
	if d[0] != "Ace of Diamonds" {
		t.Errorf("Expected first card to be 'Ace of Diamonds', but got %s", d[0])
	}
	if d[len(d)-1] != "Four of Spades" {
		t.Errorf("Expected last card to be 'Four of Spades', but got %s", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	err := d.saveToFile("_decktesting")
	if err != nil {
		t.Errorf("Error saving deck to file: %v", err)
	}
	newDeck := newDeckFromFile("_decktesting")
	if len(newDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %d", len(newDeck))
	}
}
func TestDeal(t *testing.T) {
	d := newDeck()
	hand, remainingDeck := deal(d, 5)
	if len(hand) != 5 {
		t.Errorf("Expected hand size of 5, but got %d", len(hand))
	}
	if len(remainingDeck) != 11 {
		t.Errorf("Expected remaining deck size of 11, but got %d", len(remainingDeck))
	}
}
