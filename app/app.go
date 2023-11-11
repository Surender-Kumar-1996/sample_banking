package app

import (
	"log"
	"net/http"

	"github.com/Surender-Kumar-1996/sample_banking/config"
	"github.com/Surender-Kumar-1996/sample_banking/domain"
	"github.com/Surender-Kumar-1996/sample_banking/logger"
	service "github.com/Surender-Kumar-1996/sample_banking/service"
	"github.com/gorilla/mux"
)

func Start(config *config.BankingConfig) {
	//Creating a new request Multiplexer

	router := mux.NewRouter()
	//define routes

	//wiring
	ch := CustomerHandler{service: service.NewCustomerService(domain.NewCustomerRepositoryDb(config))}
	router.HandleFunc("/customers", ch.getAllCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	//starting server
	logger.Info("Starting server on 8080")
	log.Fatalln(http.ListenAndServe(config.Server.SerAddress+config.Server.SerPort, router))
}
