package main

import (
	"customer_api_hex_arch/app"
	"log"
)

func main() {
	log.Println("Starting app on http://localhost:5555")
	app.Start()
}
