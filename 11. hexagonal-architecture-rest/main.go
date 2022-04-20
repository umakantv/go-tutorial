package main

import (
	"customer_api_hex_arch/app"
	"customer_api_hex_arch/logger"
)

func main() {
	logger.Info("Starting app on http://localhost:5555")

	app.Start()
}
