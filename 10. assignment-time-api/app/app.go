package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	r := mux.NewRouter()

	r.HandleFunc("/api/time", getCurrentTime).Methods("GET")

	http.ListenAndServe("localhost:5555", r)
}
