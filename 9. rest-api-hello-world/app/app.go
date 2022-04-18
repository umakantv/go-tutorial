package app

import (
	"log"
	"net/http"
)

func Start() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/customers", getAllCustomers)
	log.Fatal(http.ListenAndServe("localhost:5555", nil))
}
