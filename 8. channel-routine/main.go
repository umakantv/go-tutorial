package main

import (
	"fmt"
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

	for {
		go checkLink(<-c, c)
	}

}

func checkLink(link string, c chan string) {
	// This is a blocking network request
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "may be down")
		c <- link
		return
	}

	fmt.Println(link, "is up")
	c <- link
}
