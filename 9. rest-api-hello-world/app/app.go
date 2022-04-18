package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	// router := http.NewServeMux()
	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/hello", hello)
	router.HandleFunc("/customers", getAllCustomers)
	// we can also use regex in route
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer)

	log.Fatal(http.ListenAndServe("localhost:5555", router))
}
