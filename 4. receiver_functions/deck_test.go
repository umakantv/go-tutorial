package main

import (
	"os"
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

func TestSaveToFileAndLoadFromFile(t *testing.T) {
	const fileName string = "_testingFile"
	os.Remove(fileName)

	cards := createNewDeck()

	err := cards.saveToFile(fileName)

	if err != nil {
		t.Fatalf("Error in writing deck to file")
	}

	deck, error := readFromFile(fileName)

	if error != nil {
		t.Fatalf("Error in reading from file")
	}

	for i, card := range cards {
		if card != deck[i] {
			t.Fatalf("Mismatch data at file: want(%q), got(%q)", card, deck[i])
		}
	}

	os.Remove(fileName)
}
