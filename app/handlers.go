package app

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"net/http"

	service "github.com/Surender-Kumar-1996/sample_banking/Service"

)

type CustomerHandler struct {
	service service.CustomerService
}

func (ch CustomerHandler) getAllCustomer(w http.ResponseWriter, r *http.Request) {
	customers, err := ch.service.GetAllCustomer()
	if err != nil {
		log.Fatalln("Error while fetching all customers: ", err)
	}
	log.Println("Sending content type: ", r.Header.Get("Content-Type"))
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
