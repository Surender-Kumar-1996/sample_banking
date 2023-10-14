package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name    string `json:"name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zip_code"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!!")
}

func getAllCustomer(w http.ResponseWriter, r *http.Request) {
	customer := []Customer{
		{Name: "Pritam", City: "Ranchi", Zipcode: "834009"},
		{Name: "Pankaj", City: "Bhagalpur", Zipcode: "832451"},
	}
	log.Println("Sending content type: ", r.Header.Get("Content-Type"))
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customer)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}
}
