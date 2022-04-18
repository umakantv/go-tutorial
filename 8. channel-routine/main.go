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
The output at this point is the following:

----------
2022/04/18 14:25:53 Checking http://google.com
2022/04/18 14:25:53 Finished creating all routines.
2022/04/18 14:25:53 Checking http://github.com
2022/04/18 14:25:53 Checking http://amazon.com
2022/04/18 14:25:53 Checking http://stackoverflow.com
2022/04/18 14:25:53 Checking http://facebook.com
2022/04/18 14:25:53 Finished creating all routines.
2022/04/18 14:25:53 Finished creating all routines.
2022/04/18 14:25:53 Finished creating all routines.
2022/04/18 14:25:53 Finished creating all routines.
2022/04/18 14:25:53 Finished creating all routines.
2022/04/18 14:25:53 Finished creating all routines.
2022/04/18 14:25:53 Finished creating all routines.
2022/04/18 14:25:53 Finished creating all routines.
2022/04/18 14:25:53 Finished creating all routines.
2022/04/18 14:25:54 http://github.com is up
2022/04/18 14:25:54 Finished testing the link http://github.com
2022/04/18 14:25:54 Finished testing the link http://github.com
2022/04/18 14:25:54 Finished testing the link http://github.com
2022/04/18 14:25:54 Finished testing the link http://github.com
2022/04/18 14:25:54 Finished testing the link http://github.com
2022/04/18 14:25:54 Finished testing the link http://github.com
2022/04/18 14:25:54 Finished testing the link http://github.com
2022/04/18 14:25:54 Finished testing the link http://github.com
Message from channel http://github.com is up and working.
2022/04/18 14:25:54 Finished testing the link http://github.com
2022/04/18 14:25:54 Finished testing the link http://github.com
2022/04/18 14:25:54 http://google.com is up
2022/04/18 14:25:54 Finished testing the link http://google.com
2022/04/18 14:25:54 Finished testing the link http://google.com
2022/04/18 14:25:54 Finished testing the link http://google.com
2022/04/18 14:25:54 Finished testing the link http://google.com
2022/04/18 14:25:54 Finished testing the link http://google.com
2022/04/18 14:25:54 Finished testing the link http://google.com
2022/04/18 14:25:54 Finished testing the link http://google.com
2022/04/18 14:25:54 Finished testing the link http://google.com
2022/04/18 14:25:54 Finished testing the link http://google.com
2022/04/18 14:25:54 Finished testing the link http://google.com
Message from channel http://google.com is up and working.
2022/04/18 14:25:54 http://stackoverflow.com is up
2022/04/18 14:25:54 Finished testing the link http://stackoverflow.com
2022/04/18 14:25:54 Finished testing the link http://stackoverflow.com
2022/04/18 14:25:54 Finished testing the link http://stackoverflow.com
2022/04/18 14:25:54 Finished testing the link http://stackoverflow.com
2022/04/18 14:25:54 Finished testing the link http://stackoverflow.com
umakant.vashishtha@RZP2862 8. channel-routine % go run .
2022/04/18 14:30:11 Finished creating all routines.
2022/04/18 14:30:11 Finished creating all routines.
2022/04/18 14:30:11 Finished creating all routines.
2022/04/18 14:30:11 Finished creating all routines.
2022/04/18 14:30:11 Finished creating all routines.
2022/04/18 14:30:11 Finished creating all routines.
2022/04/18 14:30:11 Finished creating all routines.
2022/04/18 14:30:11 Finished creating all routines.
2022/04/18 14:30:11 Finished creating all routines.
2022/04/18 14:30:11 Finished creating all routines.
2022/04/18 14:30:11 Checking http://github.com
2022/04/18 14:30:11 Checking http://facebook.com
2022/04/18 14:30:11 Checking http://amazon.com
2022/04/18 14:30:11 Checking http://stackoverflow.com
2022/04/18 14:30:11 Checking http://google.com
2022/04/18 14:30:12 http://github.com is up
2022/04/18 14:30:12 Finished testing the link http://github.com
2022/04/18 14:30:12 Finished testing the link http://github.com
2022/04/18 14:30:12 Finished testing the link http://github.com
2022/04/18 14:30:12 Finished testing the link http://github.com
2022/04/18 14:30:12 Finished testing the link http://github.com
2022/04/18 14:30:12 Finished testing the link http://github.com
2022/04/18 14:30:12 Finished testing the link http://github.com
2022/04/18 14:30:12 Finished testing the link http://github.com
2022/04/18 14:30:12 Finished testing the link http://github.com
2022/04/18 14:30:12 Finished testing the link http://github.com
Message from channel http://github.com is up and working.
2022/04/18 14:30:12 http://google.com is up
2022/04/18 14:30:12 Finished testing the link http://google.com
2022/04/18 14:30:12 Finished testing the link http://google.com
2022/04/18 14:30:12 Finished testing the link http://google.com
2022/04/18 14:30:12 Finished testing the link http://google.com
2022/04/18 14:30:12 Finished testing the link http://google.com
2022/04/18 14:30:12 Finished testing the link http://google.com
2022/04/18 14:30:12 Finished testing the link http://google.com
2022/04/18 14:30:12 Finished testing the link http://google.com
2022/04/18 14:30:12 Finished testing the link http://google.com
2022/04/18 14:30:12 Finished testing the link http://google.com
Message from channel http://google.com is up and working.
2022/04/18 14:30:12 http://stackoverflow.com is up
2022/04/18 14:30:12 Finished testing the link http://stackoverflow.com
2022/04/18 14:30:12 Finished testing the link http://stackoverflow.com
2022/04/18 14:30:12 Finished testing the link http://stackoverflow.com
2022/04/18 14:30:12 Finished testing the link http://stackoverflow.com
2022/04/18 14:30:12 Finished testing the link http://stackoverflow.com
2022/04/18 14:30:12 Finished testing the link http://stackoverflow.com
2022/04/18 14:30:12 Finished testing the link http://stackoverflow.com
2022/04/18 14:30:12 Finished testing the link http://stackoverflow.com
2022/04/18 14:30:12 Finished testing the link http://stackoverflow.com
2022/04/18 14:30:12 Finished testing the link http://stackoverflow.com
Message from channel http://stackoverflow.com is up and working.
2022/04/18 14:30:12 http://facebook.com is up
2022/04/18 14:30:12 Finished testing the link http://facebook.com
2022/04/18 14:30:12 Finished testing the link http://facebook.com
2022/04/18 14:30:12 Finished testing the link http://facebook.com
2022/04/18 14:30:12 Finished testing the link http://facebook.com
2022/04/18 14:30:12 Finished testing the link http://facebook.com
2022/04/18 14:30:12 Finished testing the link http://facebook.com
2022/04/18 14:30:12 Finished testing the link http://facebook.com
2022/04/18 14:30:12 Finished testing the link http://facebook.com
2022/04/18 14:30:12 Finished testing the link http://facebook.com
2022/04/18 14:30:12 Finished testing the link http://facebook.com
Message from channel http://facebook.com is up and working.
2022/04/18 14:30:13 http://amazon.com is up
2022/04/18 14:30:13 Finished testing the link http://amazon.com
2022/04/18 14:30:13 Finished testing the link http://amazon.com
2022/04/18 14:30:13 Finished testing the link http://amazon.com
2022/04/18 14:30:13 Finished testing the link http://amazon.com
2022/04/18 14:30:13 Finished testing the link http://amazon.com
2022/04/18 14:30:13 Finished testing the link http://amazon.com
2022/04/18 14:30:13 Finished testing the link http://amazon.com
2022/04/18 14:30:13 Finished testing the link http://amazon.com
2022/04/18 14:30:13 Finished testing the link http://amazon.com
Message from channel http://amazon.com is up and working.
2022/04/18 14:30:13 Finished testing the link http://amazon.com
----------

Note how the line 29 is logged before logs from line 34, but that will not always be the case.
> So we can't rely on Lines 24-26 of Main routine getting executed before Lines 38 from other go routines.

But we can safely say that since all go routines will have a listener,
the program will terminate after all go routines each will have sent a message.

Interesting thing to note here is that:
And that lines 51-54 for each go routine will be executed before execution goes back to Line 32 on the Main routine.

*/
