package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fin-man/finance-manager/server/models"

	"github.com/fin-man/finance-manager/server/services"
)

type TransactionPostgresHandler struct {
	TransactionPostgresService *services.TransactionPostgresService
}

func NewTransactionPostgresHandler(transactionPostgresService *services.TransactionPostgresService) *TransactionPostgresHandler {
	return &TransactionPostgresHandler{
		TransactionPostgresService: transactionPostgresService,
	}
}

func (t *TransactionPostgresHandler) GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	//json handle

}

func (t *TransactionPostgresHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction models.TransactionModel
	err := json.NewDecoder(r.Body).Decode(&transaction)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = t.TransactionPostgresService.CreateTransaction(&transaction)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "created transaction")
}
