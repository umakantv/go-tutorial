package app

import (
	"customer_api_hex_arch/domain"
	"customer_api_hex_arch/service"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func Start() {

	// customerHandler := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	customerHandler := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDB())}
	router := mux.NewRouter()

	// routes
	router.HandleFunc("/api/customers", customerHandler.getAllCustomers)

	log.Fatal(http.ListenAndServe("localhost:5555", router))
}
