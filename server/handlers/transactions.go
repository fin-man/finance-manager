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
	enableCors(&w)

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
	tr, err := t.TransactionService.GetAllTransactions()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	enableCors(&w)

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(tr); err != nil {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

	}
}

func (t *TransactionHandler) GetAllTransactionsGraph(w http.ResponseWriter, r *http.Request) {
	testData := t.TransactionService.GetAllTransactionsGraph()
	enableCors(&w)

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(testData); err != nil {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

	}
}
func (t *TransactionHandler) GetTransactionsInDateRange(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")

	fmt.Printf("FROM : %s , TO : %s \n", from, to)
	tr, err := t.TransactionService.GetTransactionsInDateRange(from, to)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(tr); err != nil {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
