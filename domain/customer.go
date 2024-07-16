package domain

import (
	"api-banking/dto"
	"api-banking/errs"
)

// Customer represents a customer in banking
type Customer struct {
	ID          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

// CustomerRepository is a repository of customers
type CustomerRepository interface {
	FindAll(status string) (*[]Customer, *errs.AppError)
	FindByID(ID string) (*Customer, *errs.AppError)
	Transaction(transaction Transaction) *errs.AppError
}

func (c Customer) statusAsText() string {
	statusAsText := "active"
	if c.Status == "false" {
		statusAsText = "inactive"
	}

	return statusAsText
}

// ToDto ...
func (c Customer) ToDto() dto.CustomerResponse {

	return dto.CustomerResponse{

		ID:          c.ID,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.statusAsText(),
	}
}
