package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

type deck []string

// Print the cards in a deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// init sets initial values for variables used in the function.
// Go executes init functions automatically at program startup, after global variables have been initialized.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Get a newly created Deck
func createNewDeck() deck {
	newDeck := deck{}

	cardSuites := []string{
		"Spades", "Diamonds", "Clubs", "Hearts",
	}
	cardValues := []string{
		"Ace",
		"Two",
		// "Three",
		// "Four",
		// "Five",
		// "Six",
		// "Seven",
		// "Eight",
		// "Nine",
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

// Get two sets of cards from a deck
// according to a handSize
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Shuffle the deck in O(n)
func (d deck) shuffle() deck {
	l := len(d)

	for i := 0; i < l; i = i + 1 {
		position := rand.Intn(len(d))
		d[i], d[position] = d[position], d[i]
	}

	return d
}

func (d deck) toString() string {
	return strings.Join([]string(d), "\n")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func readFromFile(fileName string) (deck, error) {
	data, error := ioutil.ReadFile(fileName)
	if error != nil {
		log.Panicln("Error while reading from file: ", error.Error())
		return nil, error
	} else {
		cards := strings.Split(string(data), "\n")
		return deck(cards), nil
	}
}
