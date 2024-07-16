package service

import (
	"api-banking/domain"
	"api-banking/dto"
	"api-banking/errs"
	"time"
)

// AccountService ...
type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	GetBalance(string) (*dto.BalanceResponse, *errs.AppError)
	NewTransaction(dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError)
}

// DefaultAccountService ---
type DefaultAccountService struct {
	repo domain.AccountRepository
}

// NewAccountService ...
func NewAccountService(repository domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repository}
}

// NewAccount ...
func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	acc := domain.Account{
		AccountID:   "",
		CustomerID:  req.CustomerID,
		AccountType: req.AccountType,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		Amount:      req.Amount,
		Status:      "1",
	}

	newAccount, err := s.repo.Save(acc)
	if err != nil {
		return nil, err
	}

	response := newAccount.ToNewAccountDto()

	return &response, nil
}

// GetBalance ...
func (s DefaultAccountService) GetBalance(accontID string) (*dto.BalanceResponse, *errs.AppError) {
	var balance dto.BalanceResponse

	accountAmount, err := s.repo.Balance(accontID)
	if err != nil {
		return nil, err
	}
	balance.Amount = accountAmount

	return &balance, nil
}

// NewTransaction ..
func (s DefaultAccountService) NewTransaction(req dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	transaction := domain.Transaction{
		TransactionID:   "",
		Amount:          req.Amount,
		AccountID:       req.AccountID,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}

	var transResp *domain.Transaction
	var appErr *errs.AppError

	switch req.TransactionType {
	case "deposit":
		transResp, appErr = s.repo.Deposit(transaction)
	case "withdrawal":
		transResp, appErr = s.repo.Withdrawal(transaction)
	default:
		return nil, errs.NewValidationError("Transaction amount must be greater than 0")
	}
	if appErr != nil {
		return nil, appErr
	}

	response := dto.NewTransactionResponse{
		TransactionID:  transResp.TransactionID,
		AccountBalance: transResp.Balance,
	}

	return &response, nil
}
