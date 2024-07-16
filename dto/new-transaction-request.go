package dto

import "api-banking/errs"

// NewTransactionRequest ...
type NewTransactionRequest struct {
	AccountID       string  `json:"account_id"`
	CustomerID      string  `json:"customer_id"`
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"amount"`
}

// Validate ...
func (r NewTransactionRequest) Validate() *errs.AppError {
	if r.Amount < 0 {
		return errs.NewValidationError("Transaction amount must be greater than 0")
	}

	if r.TransactionType != "deposit" && r.TransactionType != "withdrawal" {
		return errs.NewValidationError("Transaction type should be deposit or withdrawal")
	}

	return nil
}
