package app

import (
	"log"
	"net/http"
)

func Start() {

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/customers", getAllCustomers)
	log.Fatal(http.ListenAndServe("localhost:5555", mux))
}
