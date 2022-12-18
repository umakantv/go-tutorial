package main

import (
	"customer-account-service/app"

	"github.com/umakantv/go-utils/logger"
)

func main() {
	logger.Info("Starting app on http://localhost:5555")

	app.Start()
}
