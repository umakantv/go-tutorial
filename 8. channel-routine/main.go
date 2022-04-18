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
	log.Println("Finished creating all routines")

	// This is a blocking request for a channel
	// It will wait till it receives one message from the channel
	// If there is no message, it will block the program
	fmt.Println("Message from channel", <-c)

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
	c <- link + " is up"
}

/*
The output at this point is the following:

----------

2022/04/18 13:30:45 Finished creating all routines
2022/04/18 13:30:45 Checking http://google.com
2022/04/18 13:30:45 Checking http://stackoverflow.com
2022/04/18 13:30:45 Checking http://facebook.com
2022/04/18 13:30:45 Checking http://amazon.com
2022/04/18 13:30:45 Checking http://github.com
2022/04/18 13:30:46 http://google.com is up
Message from channel http://google.com is up

----------

Note how the line 29 is logged before logs from line 34.

Also check the logs after commenting the line 29

----------

2022/04/18 13:33:01 Finished creating all routines

----------
*/
