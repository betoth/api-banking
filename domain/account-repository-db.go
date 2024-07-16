package domain

import (
	"api-banking/errs"
	"api-banking/logger"
	"database/sql"
	"strconv"

	"github.com/jmoiron/sqlx"
)

// AccountRepositoryDB ...
type AccountRepositoryDB struct {
	db *sqlx.DB
}

// NewAccountRepositoryDB ...
func NewAccountRepositoryDB(dbClient *sqlx.DB) AccountRepositoryDB {

	return AccountRepositoryDB{dbClient}
}

// Save ...
func (d AccountRepositoryDB) Save(acc Account) (*Account, *errs.AppError) {
	insertAccountSQL := `
	INSERT INTO accounts
	(customer_id, opening_date, account_type, amount, status)
	VALUES($1, $2, $3, $4, $5)
	RETURNING account_id;`

	var accountID int64
	err := d.db.QueryRow(insertAccountSQL, acc.CustomerID, acc.OpeningDate, acc.AccountType, acc.Amount, acc.Status).Scan(&accountID)
	if err != nil {
		logger.Error("Error creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	acc.AccountID = strconv.FormatInt(accountID, 10)

	return &acc, nil
}

// Balance ...
func (d AccountRepositoryDB) Balance(accountID string) (float64, *errs.AppError) {
	var amount float64
	selectBalanceSQL := "SELECT amount FROM accounts where account_id = $1"

	err := d.db.Get(&amount, selectBalanceSQL, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Info("No rows returs")
			return 0, errs.NewNotFoundError("Account not found")
		}
		logger.Error("error searching for balance:" + err.Error())
		return 0, errs.NewUnexpectedError("Unexpected database error")
	}

	return amount, nil
}

// Deposit ...
func (d AccountRepositoryDB) Deposit(transaction Transaction) (*Transaction, *errs.AppError) {

	tx, err := d.db.Begin()
	if err != nil {
		logger.Error("Error creating tx: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	insertTransactionSQL := `INSERT INTO public.transactions
	(account_id, amount, transaction_type, transaction_date)
	VALUES($1, $2, 'deposit', $3)
	RETURNING transaction_id;`
	var transactionID string
	err = d.db.QueryRow(insertTransactionSQL, transaction.AccountID, transaction.Amount, transaction.TransactionDate).Scan(&transactionID)
	if err != nil {
		tx.Rollback()
		logger.Error("Error creating new transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	updateAmountSQL := "UPDATE accounts SET amount = amount + $1 where account_id = $2 RETURNING amount;"
	var balance float64
	err = d.db.QueryRow(updateAmountSQL, transaction.Amount, transaction.AccountID).Scan(&balance)
	if err != nil {
		tx.Rollback()
		logger.Error("Error creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	err = tx.Commit()
	if err != nil {
		logger.Error("Error commit tx: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	transaction.Balance = balance
	transaction.TransactionID = transactionID
	return &transaction, nil
}

// Withdrawal ...
func (d AccountRepositoryDB) Withdrawal(transaction Transaction) (*Transaction, *errs.AppError) {

	accountBalance, appErr := d.Balance(transaction.AccountID)
	if appErr != nil {
		return nil, appErr
	}

	if accountBalance < transaction.Amount {
		return nil, errs.NewValidationError("Account does not have sufficient balance")
	}

	tx, err := d.db.Begin()
	if err != nil {
		logger.Error("Error creating tx: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	insertTransactionSQL := `INSERT INTO public.transactions
	(account_id, amount, transaction_type, transaction_date)
	VALUES($1, $2, 'withdrawal', $3)
	RETURNING transaction_id;`
	var transactionID string
	err = d.db.QueryRow(insertTransactionSQL, transaction.AccountID, transaction.Amount, transaction.TransactionDate).Scan(&transactionID)
	if err != nil {
		tx.Rollback()
		logger.Error("Error creating new transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	updateAmountSQL := "UPDATE accounts SET amount = amount - $1 where account_id = $2 RETURNING amount;"
	var balance float64
	err = d.db.QueryRow(updateAmountSQL, transaction.Amount, transaction.AccountID).Scan(&balance)
	if err != nil {
		tx.Rollback()
		logger.Error("Error creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	err = tx.Commit()
	if err != nil {
		logger.Error("Error commit tx: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	transaction.Balance = balance
	transaction.TransactionID = transactionID
	return &transaction, nil
}
