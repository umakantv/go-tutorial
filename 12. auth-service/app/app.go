package app

import (
	"auth/domain"
	"auth/handlers"
	"auth/service"
	"fmt"
	"log"
	"net/http"

	"github.com/umakantv/go-utils/db"

	"github.com/umakantv/go-utils/logger"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

var config Config

func init() {
	LoadConfig(&config)
	logger.Init(config.Logger)
}

func Start() {

	dbClient := db.GetDBConnection(config.Database)
	router := mux.NewRouter()
	authRepository := domain.NewAuthRepository(dbClient)
	ah := handlers.AuthHandler{
		Service: service.NewLoginService(authRepository, domain.GetRolePermissions()),
	}

	router.HandleFunc("/auth/login", ah.Login).Methods(http.MethodPost)
	router.HandleFunc("/auth/register", ah.NotImplementedHandler).Methods(http.MethodPost)
	router.HandleFunc("/auth/refresh", ah.Refresh).Methods(http.MethodPost)
	router.HandleFunc("/auth/verify", ah.Verify).Methods(http.MethodGet)

	address := config.App.HOST
	port := config.App.PORT

	logger.Info(fmt.Sprintf("Starting OAuth server on %s:%s ...", address, port))

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}
