package handlers

import (
	"encoding/json"
	"finance-manager/categories"
	"finance-manager/server/services"
	"finance-manager/utils"
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

	transaction.MakeAmountPositive() //some banks have a mount as negative
	transaction.ToLowerCase()

	data, err := json.Marshal(transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	generatedID := t.generateID(transaction.Bank, transaction.Amount, transaction.Category, transaction.TransactionDate, transaction.Description)
	err = t.TransactionService.CreateTransaction(data, generatedID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responseMessage := fmt.Sprintf(CreatedResponse, string(data), "created_successfully")
	fmt.Fprintf(w, responseMessage)
}

func (t *TransactionHandler) SearchTransactions(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	search := r.URL.Query().Get("search")
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")

	tr, err := t.TransactionService.SearchTransaction(search, from, to)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(tr); err != nil {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

	}
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

	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")

	testData := t.TransactionService.GetAllTransactionsGraph(from, to)
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

func (t *TransactionHandler) generateID(bank categories.Bank, amount float64, category categories.Category, date, description string) string {
	return utils.EncodeToBase64(fmt.Sprintf("%s-%f-%s-%s-%s", bank, amount, category, date, description))
}
