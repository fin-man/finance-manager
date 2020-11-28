package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/fin-man/finance-manager/server/models"

	"github.com/fin-man/finance-manager/categories"
	"github.com/fin-man/finance-manager/server/services"
	"github.com/jinzhu/gorm"
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
	startDate, okStartDate := r.URL.Query()["startdate"]
	endDate, okEndDate := r.URL.Query()["enddate"]
	query := make(map[string][]string)

	if okBanks {
		query["bank"] = banks
	}

	if okCat {
		query["category"] = categories
	}

	if !okStartDate || !okEndDate || len(startDate) != 1 || len(endDate) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//time format validations
	//if all goes well

	transactions, err := t.TransactionPostgresService.SearchTransactions(query)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := make(map[string]interface{})

	resp["transactions"] = transactions
	resp["size"] = len(transactions)

	jsonResp, err := json.Marshal(resp)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)

}

func (t *TransactionPostgresHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	type bindTransactionModel struct {
		gorm.Model
		TransactionID   string              `json:"transactions_id" gorm:"unique"`
		TransactionDate string              `csv:"transaction_date" json:"transaction_date"`
		Amount          float64             `csv:"amount"  json:"amount"`
		Description     string              `csv:"description"  json:"description"`
		Bank            categories.Bank     `csv:"bank"  json:"bank"`
		AccountID       string              `csv:"account_id"  json:"account_id"`
		Category        categories.Category `csv:"category"  json:"category"`
		AccountType     string              `csv:"account_type" json:"account_type`
		Hash            string              `csv:"-" json:"hash"`
	}

	var bindTransaction bindTransactionModel
	err := json.NewDecoder(r.Body).Decode(&bindTransaction)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	convertedTime, err := time.Parse("2006-01-02", bindTransaction.TransactionDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//reapply to a transaction model
	var transaction models.TransactionModel
	transaction.TransactionID = bindTransaction.TransactionID
	transaction.TransactionDate = convertedTime
	transaction.Amount = bindTransaction.Amount
	transaction.Description = bindTransaction.Description
	transaction.Bank = bindTransaction.Bank
	transaction.AccountID = bindTransaction.AccountID
	transaction.Category = bindTransaction.Category
	transaction.AccountType = bindTransaction.AccountType
	transaction.Hash = bindTransaction.Hash

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
