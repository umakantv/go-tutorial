package main

import (
	"fmt"
	"math/rand"
	"time"
)

type bot interface {
	getGreeting() string
}

type enhancedBot interface {
	getRandomGreeting() string
}

// Implements `bot` interface
type englishBot struct{}

// Implements `bot` interface implicitly
type spanishBot struct{}

func init() {
	rand.Seed(time.Now().UnixMicro())
}

func (englishBot) getRandomGreeting() string {
	greetings := []string{
		"Hi there.",
		"Hello, mate.",
		"What's up?",
	}

	return greetings[rand.Intn(len(greetings))]
}

func (englishBot) getGreeting() string {
	return "Hi there"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func getGreeting(b bot) {

	/*
		// Checking if an object from one interface responds to a method in other interface
		// In other words, we can check if an instance a from interface I1 also satisfies another interface I2
		// https://stackoverflow.com/questions/29684609/how-to-check-if-an-object-has-a-particular-method

	*/

	if ebWithAssert, ok := b.(enhancedBot); ok {
		// We can also do the above check with one declaration
		// if ebWithAssert, ok := b.(interface{ getRandomGreeting() string }); ok {
		fmt.Println(ebWithAssert.getRandomGreeting())
		return
	}

	fmt.Println(b.getGreeting())
}

func Bot() {
	eb := englishBot{}
	sb := spanishBot{}

	getGreeting(eb)
	getGreeting(sb)

}
