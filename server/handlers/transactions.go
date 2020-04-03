package handlers

import (
	"encoding/json"
	"finance-manager/categories"
	"finance-manager/server/services"
	"fmt"
	"net/http"
)

type TransactionHandler struct {
	TransactionService *services.TransactionService
}

var (
	CreatedResponse string = `{ "transaction" : %s , "message" : %s }`
)

func NewTransactionHandler() *TransactionHandler {

	transactionService := services.NewTransactionService()
	return &TransactionHandler{
		TransactionService: transactionService,
	}
}

func (t *TransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction categories.NormalizedTransaction

	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := json.Marshal(transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = t.TransactionService.CreateTransaction(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responseMessage := fmt.Sprintf(CreatedResponse, string(data), "created_successfully")
	fmt.Fprintf(w, responseMessage)
}

func (t *TransactionHandler) GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	t.TransactionService.GetAllTransactions()
}
