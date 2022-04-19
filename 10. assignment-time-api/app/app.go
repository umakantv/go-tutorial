package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	r := mux.NewRouter()

	r.HandleFunc("/api/time", getCurrentTime).Methods("GET")

	log.Fatal(http.ListenAndServe("localhost:5555", r))

	log.Println("Server is ready at http://localhost:5555")
}
