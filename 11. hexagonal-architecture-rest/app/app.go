package app

import (
	"customer_api_hex_arch/domain"
	"customer_api_hex_arch/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	customerHandler := CustomerHandlers{
		service.NewCustomerService(domain.NewCustomerRepositoryStub()),
	}
	router := mux.NewRouter()

	// routes
	router.HandleFunc("/api/customers", customerHandler.getAllCustomers)

	log.Fatal(http.ListenAndServe("localhost:5555", router))
}
