package main

import (
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

	for _, link := range links {
		go checkLink(link)
	}

}

func checkLink(link string) {
	_, err := http.Get(link)

	if err != nil {
		log.Println(link, "may be down")
		return
	}

	log.Println(link, "is up")
}
