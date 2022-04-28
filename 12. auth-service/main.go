package main

import (
	"auth/app"
	"auth/logger"
)

func main() {
	logger.Info("Starting app on http://localhost:5555")

	app.Start()
}
