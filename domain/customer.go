package domain

import (
	"github.com/Surender-Kumar-1996/sample_banking/dto"
	"github.com/Surender-Kumar-1996/sample_banking/errs"
)

// Business Side logic
type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

func (c Customer) statusAsText() string {
	statusAsText := "active"

	if c.Status == "0" {
		statusAsText = "inactive"
	}

	return statusAsText
}

func (c Customer) ToDto() *dto.CustomerResponse {

	return &dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.statusAsText(),
	}
}

// Secondary Port
type CustomerRepository interface {
	FindAll(string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
