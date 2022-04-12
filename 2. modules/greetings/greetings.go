package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), name)
	// message := fmt.Sprintf(randomFormat()) // this breaks the test
	return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	result := make(map[string]string)

	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	for i := 0; i < len(names); i++ {
		message, error := Hello(names[i])
		if error != nil {
			return nil, error
		}
		result[names[i]] = message
	}

	/*
		Another for loop:
		for index, value := range names {
		}

		For variables we don't need we use _ (the blank identifier)
	*/
	return result, nil
}

// init sets initial values for variables used in the function.
// Go executes init functions automatically at program startup, after global variables have been initialized.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
		"Ahoy sexy, %v!!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
