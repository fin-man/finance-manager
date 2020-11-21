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
	banks, okBanks := r.URL.Query()["banks"]
	categories, okCat := r.URL.Query()["categories"]
	query := make(map[string][]string)

	if okBanks {
		query["bank"] = banks
	}

	if okCat {
		query["category"] = categories
	}

	transactions, err := t.TransactionPostgresService.SearchTransactions(query)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonResp, err := json.Marshal(transactions)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)

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

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
