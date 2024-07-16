package domain

import (
	"api-banking/dto"
	"api-banking/errs"
)

// Account ...
type Account struct {
	AccountID   string
	CustomerID  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

// AccountRepository ...
type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	Balance(AccountID string) (float64, *errs.AppError)
	Deposit(transaction Transaction) (*Transaction, *errs.AppError)
	Withdrawal(transaction Transaction) (*Transaction, *errs.AppError)
}

// ToNewAccountDto ...
func (a Account) ToNewAccountDto() dto.NewAccountResponse {

	return dto.NewAccountResponse{
		AccountID: a.AccountID,
	}

}
