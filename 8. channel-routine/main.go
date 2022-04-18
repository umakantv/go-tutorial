package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://amazon.com",
		"http://github.com",
	}

	c := make(chan string)

	for _, link := range links {
		// Create child go routines
		go checkLink(link, c)
	}
	for i := 0; i < 10; i++ {
		log.Println(`Finished creating all routines.`)
	}

	for i := 0; i < len(links); i++ {
		// This is a blocking request for a channel
		// It will wait till it receives one message from the channel
		// If there is no message, it will block the program
		fmt.Println("Message from channel", <-c)
	}

}

func checkLink(link string, c chan string) {
	log.Println("Checking", link)
	// This is a blocking network request
	_, err := http.Get(link)

	if err != nil {
		log.Println(link, "may be down")
		c <- link + " may be down"
		return
	}

	log.Println(link, "is up")
	c <- link + " is up and working."

	for i := 0; i < 10; i++ {
		// log.Println(`Finished creating all routines.`)
		log.Println("Finished testing the link", link)
	}
}

/*


We can't rely on Lines 24-26 of Main routine getting executed before Lines 38 from other go routines.
They can get executed in any order.

Similarly:
And that lines 51-54 for each go routine will not always be executed before
the execution goes back to Line 32 on the Main routine.

But we can safely say that since all go routines will have a listener,
the program will terminate after all go routines each will have sent a message.

*/
