package main

import (
	"log"
	"net/http"
	"time"
)

var links = []string{
	"http://google.com",
	"http://facebook.com",
	"http://stackoverflow.com",
	"http://amazon.com",
	"http://github.com",
}

func checkSites() string {

	c := make(chan string)

	for _, link := range links {
		// Create child go routines
		go checkLink(link, c)
	}
	duration := time.Second * 5

	// Ranging over channel is a blocking operation
	// It will wait for the channel to return a value
	for l := range c {
		// fmt.Println(l)
		go func(l string) {
			log.Println("Checked from site_status - scheduling next check in 5 seconds", l)
			time.Sleep(duration)
			checkLink(l, c)
		}(l)
	}

	return "Done"
}

func checkLink(link string, c chan string) {
	// This is a blocking network request
	_, err := http.Get(link)

	if err != nil {
		log.Println(link, "may be down")
		c <- link
		return
	}

	log.Println(link, "is up")
	c <- link
}

func Main() {

	checkSites()
	// multipleChannels()

}
