package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	// router := http.NewServeMux()
	router := mux.NewRouter()
	router.HandleFunc("/hello", hello)
	router.HandleFunc("/customers", getAllCustomers)
	router.HandleFunc("/customers/{customer_id}", getCustomer)
	log.Fatal(http.ListenAndServe("localhost:5555", router))
}
