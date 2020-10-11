package handlers

import (
	"net/http"

	"github.com/fin-man/finance-manager/server/services"
)

type TransactionPostgresHandler struct {
	TransactionPostgresService *services.TransactionPostgresService
}

func (t *TransactionPostgresHandler) GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	//json handle
}

func (t *TransactionPostgresHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {

}
