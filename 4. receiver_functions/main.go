package main

import "fmt"

func main() {
	cards := createNewDeck()
	cards = cards.shuffle()
	fmt.Println("Total cards", len(cards))

	hand, remainingCards := deal(cards, 5)
	cards = remainingCards

	fmt.Println("Hand", len(hand))
	hand.print()
	fmt.Println("Remaining Deck", len(cards))
	cards.print()

}
