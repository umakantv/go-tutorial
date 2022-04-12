package main

import (
	"fmt"
	"log"
)

func main() {
	cards := createNewDeck()
	cards = cards.shuffle()

	const fileNmae string = "initial_deck.txt"

	error := cards.saveToFile(fileNmae)
	if error != nil {
		log.Fatalln("Error in saving deck to file:", error.Error())
	} else {
		log.Println("Saved deck to file named initial_deck!")
	}

	cards, error = readFromFile(fileNmae)
	if error != nil {
		log.Panicln("Error in reading deck from file:", error.Error())
	} else {
		log.Println("Loaded deck to file named initial_deck!")
	}

	fmt.Println("Total cards", len(cards))

	hand, remainingCards := deal(cards, 5)
	cards = remainingCards

	fmt.Println("Hand", len(hand))
	hand.print()
	fmt.Println("Remaining Deck", len(cards))
	cards.print()

}
