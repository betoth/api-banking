package service

import (
	"api-banking/domain"
	"api-banking/dto"
	"api-banking/errs"
)

// CustomerService ...
type CustomerService interface {
	GetAllCustomers(status string) (*[]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

// DefaultCustomerService ...
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// NewCustomerService ...
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}

// GetAllCustomers ...
func (s DefaultCustomerService) GetAllCustomers(status string) (*[]dto.CustomerResponse, *errs.AppError) {
	switch status {
	case "active":
		status = "true"

	case "inactive":
		status = "false"

	default:
		status = ""
	}
	cs, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}
	var response []dto.CustomerResponse

	for _, c := range *cs {
		response = append(response, c.ToDto())
	}

	return &response, nil
}

// GetCustomer ...
func (s DefaultCustomerService) GetCustomer(ID string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.FindByID(ID)
	if err != nil {
		return nil, err
	}

	response := c.ToDto()

	return &response, nil
}
