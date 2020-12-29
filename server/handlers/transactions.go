package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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
	startDate, okStartDate := r.URL.Query()["startdate"]
	endDate, okEndDate := r.URL.Query()["enddate"]
	categoryFormat, okCategoryFormat := r.URL.Query()["categoryformat"]
	query := make(map[string][]string)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

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
	startDateTimeTime, err := time.Parse("2006-01-02", startDate[0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	endDateTimeTime, err := time.Parse("2006-01-02", endDate[0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	transactions, err := t.TransactionPostgresService.SearchTransactions(query, startDateTimeTime, endDateTimeTime)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if okCategoryFormat {
		if len(categoryFormat) == 1 && categoryFormat[0] != "" {
			categoryData := make(map[string][]models.TransactionModel)

			for _, transaction := range transactions {

				categoryData[string(transaction.Category)] = append(categoryData[string(transaction.Category)], transaction)
			}

			jsonResp, err := json.Marshal(categoryData)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write(jsonResp)
			return
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}

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

	var transaction models.TransactionModel
	err := json.NewDecoder(r.Body).Decode(&transaction)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// convertedTime, err := time.Parse("2006-01-02", bindTransaction.TransactionDate)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// //reapply to a transaction model
	// transaction.TransactionID = bindTransaction.TransactionID
	// transaction.TransactionDate = convertedTime
	// transaction.Amount = bindTransaction.Amount
	// transaction.Description = bindTransaction.Description
	// transaction.Bank = bindTransaction.Bank
	// transaction.AccountID = bindTransaction.AccountID
	// transaction.Category = bindTransaction.Category
	// transaction.AccountType = bindTransaction.AccountType
	// transaction.Hash = bindTransaction.Hash

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
