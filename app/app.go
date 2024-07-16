package app

import (
	"api-banking/config"
	"api-banking/domain"
	"api-banking/logger"
	"api-banking/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

// Start starts a http server
func Start(cfg *config.Config) {

	router := mux.NewRouter()
	dbClient := getDbClient(cfg)

	customerRepositoryDB := domain.NewCustomerRepositoryDB(dbClient)
	accountRepositoryDB := domain.NewAccountRepositoryDB(dbClient)

	ch := CustomerHandlers{service.NewCustomerService(customerRepositoryDB)}
	ah := AccountHanddlers{service.NewAccountService(accountRepositoryDB)}

	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.GetCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)
	router.HandleFunc("/account/{account_id:[0-9]+}/balance", ah.Balance).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}/transaction", ah.NewTransaction).Methods(http.MethodPost)

	serverConfig := fmt.Sprintf("%s:%s", cfg.ServerAddress, cfg.ServerPort)

	logger.Info("Application listening on " + serverConfig)
	log.Fatal(http.ListenAndServe(serverConfig, router))
}

func getDbClient(cfg *config.Config) *sqlx.DB {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode)
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
