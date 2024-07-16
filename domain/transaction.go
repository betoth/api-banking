package domain

// Transaction ...
type Transaction struct {
	TransactionID   string
	AccountID       string
	Amount          float64
	TransactionType string
	TransactionDate string
	Balance         float64
}
