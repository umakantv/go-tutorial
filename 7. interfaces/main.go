package main

import "fmt"

type bot interface {
	getGreeting() string
}

// Implements `bot` interface
type englishBot struct{}

// Implements `bot` interface
type spanishBot struct{}

func (englishBot) getGreeting() string {
	return "Hi there"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func getGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	getGreeting(eb)
	getGreeting(sb)
}
