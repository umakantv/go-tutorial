package main

import (
	"auth/app"

	"github.com/umakantv/go-utils/logger"
)

func main() {
	logger.Info("Starting app on the given port")

	app.Start()
}
