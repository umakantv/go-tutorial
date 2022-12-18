package main

import (
	"auth/app"

	"github.com/umakantv/go-utils/logger"
)

func main() {
	logger.Info("Starting app on http://localhost:5555")

	app.Start()
}
