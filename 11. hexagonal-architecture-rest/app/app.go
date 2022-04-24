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

	dbClient := domain.GetDBConnection()
	customerRepo := domain.NewCustomerRepositoryDB(dbClient)
	accountRepo := domain.NewAccountRepositoryDB(dbClient)

	customerHandler := CustomerHandlers{service.NewCustomerService(customerRepo)}
	accountHandlers := AccountHandlers{service.NewAccountService(accountRepo)}

	router := mux.NewRouter()

	// routes
	router.HandleFunc("/api/customers", customerHandler.getAllCustomers)
	router.HandleFunc("/api/customers/{customer_id:[0-9]+}", customerHandler.GetCustomerById)

	router.HandleFunc("/api/customers/{customer_id:[0-9]+}/account", accountHandlers.createNewAccount).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe("localhost:5555", router))
}
