package main

import (
	"greetings"
	"log"

	"github.com/razorpay/goutils/uniqueid"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(3)

	// A slice of names.
	names := []string{"Umakant", "Anand", "Vivek"}
	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Println("Error in greetings:", err)
		// log.Fatal(err)
	}
	// If no error was returned, print the returned map of
	// messages to the console.
	// fmt.Println(messages)
	for name, greeting := range messages {
		id, err := uniqueid.New()
		if err != nil {
			log.Fatalf(err.Error())
		}
		log.Println(id, name, "=>", greeting)
	}

	// Request a greeting message.
	message, err := greetings.Hello("Umakant")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Println("Error in the greeting:", err)
	}

	log.Println("Received Message\n", message, "\nEnd of message")
}
