package main

import (
	"regexp"
	"testing"
)

func TestCreateNewDeck(t *testing.T) {
	deck := createNewDeck()
	if len(deck) < 1 {
		t.Fail()
	}
	pattern := regexp.MustCompile(`\b` + " of " + `\b`)

	for _, card := range deck {
		if card == "" {
			t.Fatalf("Empty card!")
		}
		if !pattern.Match([]byte(card)) {
			t.Fatalf("Card value(%q) does not match the desired pattern %#q", card, pattern)
		}
	}
}
