package handlers

import (
	"finance-manager/server/services"
	"net/http"
)

type TransactionHandler struct {
	TransactionService *services.TransactionService
}

func NewTransactionHandler() *TransactionHandler {

	transactionService := services.NewTransactionService()
	return &TransactionHandler{
		TransactionService: transactionService,
	}
}

func (t *TransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	t.TransactionService.CreateTransaction()
}

func (t *TransactionHandler) GetAllTransactions(w http.ResponseWriter, r *http.Request) {

}
