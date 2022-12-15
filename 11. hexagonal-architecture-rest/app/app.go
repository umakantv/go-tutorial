package app

import (
	"customer_api_hex_arch/domain"
	"customer_api_hex_arch/handlers"
	"customer_api_hex_arch/logger"
	"customer_api_hex_arch/service"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var config Config

func init() {
	LoadConfig(&config)
	logger.Init(config.Logger)
}

func Start() {

	dbClient := domain.GetDBConnection(config.Database)
	customerRepo := domain.NewCustomerRepositoryDB(dbClient)
	accountRepo := domain.NewAccountRepositoryDB(dbClient)
	transactionRepo := domain.NewTransactionRepositoryDB(dbClient)

	ch := handlers.CustomerHandlers{
		Service: service.NewCustomerService(customerRepo),
	}
	ah := handlers.AccountHandlers{
		Service: service.NewAccountService(accountRepo),
	}
	th := handlers.TransactionHandler{
		Service: service.NewTransactionService(transactionRepo, accountRepo),
	}

	router := mux.NewRouter()

	// routes
	router.HandleFunc("/api/customers", ch.GetAllCustomers)
	router.HandleFunc("/api/customers/{customer_id:[0-9]+}", ch.GetCustomerById)

	router.HandleFunc("/api/customers/{customer_id:[0-9]+}/account", ah.CreateNewAccount).Methods(http.MethodPost)
	router.HandleFunc("/api/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}/transaction", th.AddNewTransaction).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe("localhost:5555", router))
}
