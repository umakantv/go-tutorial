package app

import (
	"auth/logger"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var config Config

func init() {
	LoadConfig(&config)
	logger.Init(config.Logger)
}

func Start() {

	// dbClient := db.GetDBConnection(config.Database)

	router := mux.NewRouter()

	// routes

	log.Fatal(http.ListenAndServe("localhost:5555", router))
}
