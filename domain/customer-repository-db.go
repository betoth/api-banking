package domain

import (
	"api-banking/errs"
	"api-banking/logger"
	"database/sql"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //DB
)

// CustomerRepositoryDB ...
type CustomerRepositoryDB struct {
	db *sqlx.DB
}

// NewCustomerRepositoryDB ...
func NewCustomerRepositoryDB(dbClient *sqlx.DB) CustomerRepositoryDB {

	return CustomerRepositoryDB{dbClient}
}

// FindAll ...
func (d CustomerRepositoryDB) FindAll(status string) (*[]Customer, *errs.AppError) {
	var err error
	var customers []Customer

	if status != "" {
		findAllSQL := "SELECT customer_id, name, date_of_birth, city, zipcode, status FROM customers WHERE status = $1"
		err = d.db.Select(&customers, findAllSQL, status)
	} else {
		findAllSQL := "SELECT customer_id, name, date_of_birth, city, zipcode, status FROM customers"
		err = d.db.Select(&customers, findAllSQL)
	}

	if err != nil {
		logger.Error("Error executing query " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")

	}

	return &customers, nil
}

// FindByID ..
func (d CustomerRepositoryDB) FindByID(ID string) (*Customer, *errs.AppError) {
	var customer Customer

	selectByID := "Select * from customers where customer_id = $1"
	err := d.db.Get(&customer, selectByID, ID)

	if err != nil {
		if err == sql.ErrNoRows {
			logger.Info("No rows returs")
			return nil, errs.NewNotFoundError("Customer not found")
		}

		logger.Error("Error executing scan " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return &customer, nil
}

// Transaction ...
func (d CustomerRepositoryDB) Transaction(transaction Transaction) *errs.AppError {

	return nil
}
