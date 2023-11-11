package domain

//Server Side logic

import (
	"github.com/Surender-Kumar-1996/sample_banking/errs"
	_ "github.com/go-sql-driver/mysql"
)

// Mock/stub Adapter
type CustomerRepositoryStub struct {
	customer []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, *errs.AppError) {
	return s.customer, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		Customer{Id: "1001", Name: "Ashish", City: "New Delhi", Zipcode: "834009", DateOfBirth: "19/06/1996", Status: "1"},
		Customer{Id: "1002", Name: "Sumant", City: "Ranchi", Zipcode: "123456", DateOfBirth: "12/2/1945", Status: "1"},
	}
	return CustomerRepositoryStub{customer: customers}
}
