package domain

// CustomerRepositoryStub ...
type CustomerRepositoryStub struct {
	customers []Customer
}

// FindAll find all customers
func (r CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return r.customers, nil
}

// NewCustomerRepositoryStub create a new customer repository stub
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{
			ID:          "1",
			Name:        "Gilberto",
			City:        "São José",
			Zipcode:     "88010-450",
			DateOfBirth: "30/07/1986",
			Status:      "1",
		},
		{
			ID:          "2",
			Name:        "Helen",
			City:        "São José",
			Zipcode:     "88010-450",
			DateOfBirth: "30/07/1986",
			Status:      "1",
		},
		{
			ID:          "3",
			Name:        "Noeli",
			City:        "Estrela",
			Zipcode:     "88010-007",
			DateOfBirth: "30/07/1986",
			Status:      "1",
		},
	}
	return CustomerRepositoryStub{customers}
}
