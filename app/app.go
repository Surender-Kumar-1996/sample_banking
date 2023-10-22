package app

import (
	"log"
	"net/http"

	service "github.com/Surender-Kumar-1996/sample_banking/Service"
	"github.com/Surender-Kumar-1996/sample_banking/domain"
	"github.com/gorilla/mux"
)

func Start() {
	//Creating a new request Multiplexer
	router := mux.NewRouter()
	//define routes

	//wiring
	ch := CustomerHandler{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	router.HandleFunc("/customers", ch.getAllCustomer).Methods(http.MethodGet)

	//starting server
	log.Println("Starting server on :8080")
	log.Fatalln(http.ListenAndServe("localhost:8080", router))
}
