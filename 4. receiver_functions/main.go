package main

import (
	"fmt"
	"log"
)

func main() {
	cards := createNewDeck()
	cards = cards.shuffle()

	error := cards.saveToFile("initial_deck.txt")
	if error != nil {
		log.Panicln("Error in saving deck to file:", error.Error())
	} else {
		log.Println("Saved deck to file named initial_deck!")
	}
	fmt.Println("Total cards", len(cards))

	hand, remainingCards := deal(cards, 5)
	cards = remainingCards

	fmt.Println("Hand", len(hand))
	hand.print()
	fmt.Println("Remaining Deck", len(cards))
	cards.print()

}
