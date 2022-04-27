package app

import (
	"customer_api_hex_arch/domain"
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

	ch := CustomerHandlers{service.NewCustomerService(customerRepo)}
	ah := AccountHandlers{service.NewAccountService(accountRepo)}
	th := TransactionHandler{service.NewTransactionService(transactionRepo, accountRepo)}

	router := mux.NewRouter()

	// routes
	router.HandleFunc("/api/customers", ch.getAllCustomers)
	router.HandleFunc("/api/customers/{customer_id:[0-9]+}", ch.GetCustomerById)

	router.HandleFunc("/api/customers/{customer_id:[0-9]+}/account", ah.createNewAccount).Methods(http.MethodPost)
	router.HandleFunc("/api/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}/transaction", th.addNewTransaction).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe("localhost:5555", router))
}
