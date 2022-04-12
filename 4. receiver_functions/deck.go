package main

import "fmt"

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func createNewDeck() deck {
	newDeck := deck{}

	cardSuites := []string{"Spades", "Diamonds", "Clubs", "Hearts"}
	cardValues := []string{
		"Ace",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Joker",
		"Queen",
		"King",
	}

	for _, suite := range cardSuites {
		for _, value := range cardValues {
			newDeck = append(newDeck, value+" of "+suite)
		}
	}

	return newDeck
}
