package app

import (
	"api-banking/dto"
	"api-banking/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// AccountHanddlers ...
type AccountHanddlers struct {
	service service.AccountService
}

// NewAccount ...
func (a *AccountHanddlers) NewAccount(w http.ResponseWriter, r *http.Request) {
	var request dto.NewAccountRequest

	vars := mux.Vars(r)

	customerID := vars["customer_id"]

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	request.CustomerID = customerID

	account, appError := a.service.NewAccount(request)
	if appError != nil {
		writeResponse(w, appError.Code, appError.Message)
		return
	}
	writeResponse(w, http.StatusCreated, account)
}

// Balance ...
func (a *AccountHanddlers) Balance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	accountID := vars["account_id"]

	balance, appError := a.service.GetBalance(accountID)
	if appError != nil {
		writeResponse(w, appError.Code, appError.Message)
		return
	}

	writeResponse(w, http.StatusCreated, balance)

}

// NewTransaction ...
func (a *AccountHanddlers) NewTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction dto.NewTransactionRequest
	vars := mux.Vars(r)
	accountID := vars["account_id"]
	customerID := vars["customer_id"]

	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err)
		return
	}
	transaction.AccountID = accountID
	transaction.CustomerID = customerID

	responseTransaction, appErr := a.service.NewTransaction(transaction)
	if appErr != nil {
		writeResponse(w, appErr.Code, appErr.Message)
		return
	}

	writeResponse(w, http.StatusCreated, responseTransaction)

}
