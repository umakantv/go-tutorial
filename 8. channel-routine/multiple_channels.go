package main

import "log"

func MultipleChannels() {

	for _, l := range links {
		ch := make(chan string)
		go checkLink(l, ch)
		l := <-ch

		log.Println("Checked from multiple channels", l)
	}
}
