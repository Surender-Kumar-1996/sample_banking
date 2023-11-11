package service

import (
	"github.com/Surender-Kumar-1996/sample_banking/domain"
	"github.com/Surender-Kumar-1996/sample_banking/dto"
	"github.com/Surender-Kumar-1996/sample_banking/errs"
)

// service interface
type CustomerService interface {
	GetAllCustomer(string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]dto.CustomerResponse, *errs.AppError) {
	if status == "inactive" {
		status = "0"
	} else if status == "active" {
		status = "1"
	} else {
		status = ""
	}
	c, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}
	//map customer to customer response using DTO
	cr := make([]dto.CustomerResponse, 0)
	for _, v := range c {
		cr = append(cr, *v.ToDto())
	}
	return cr, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	return c.ToDto(), nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
