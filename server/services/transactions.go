package services

import (
	"fmt"
	"strings"
	"time"

	"github.com/fin-man/finance-manager/categories"
	"github.com/fin-man/finance-manager/server/models"
	"github.com/fin-man/finance-manager/utils"
)

type TransactionPostgresService struct {
	TransactionModel *models.TransactionModel
}

func NewTransactionPostgresService() *TransactionPostgresService {
	return &TransactionPostgresService{}
}

func (t *TransactionPostgresService) GetAllTransactions() ([]models.TransactionModel, error) {
	return t.TransactionModel.GetAllTransactions()
}

func (t *TransactionPostgresService) CreateTransaction(transaction *models.TransactionModel) (*models.TransactionModel, error) {
	id := t.generateID(transaction.Bank, transaction.Amount, transaction.Category, transaction.TransactionDate, transaction.Description)
	transaction.Hash = id
	return t.TransactionModel.CreateTransaction(transaction)
}

func (t *TransactionPostgresService) generateID(bank categories.Bank, amount float64, category categories.Category, date int64, description string) string {
	return utils.EncodeToBase64(fmt.Sprintf("%s-%f-%s-%d-%s", bank, amount, category, date, description))
}
func (t *TransactionPostgresService) SearchTransactions(query map[string][]string, startTime time.Time, endTime time.Time) ([]models.TransactionModel, error) {

	transactions, err := t.TransactionModel.SearchTransactions(query, startTime.Unix(), endTime.Unix())

	if err != nil {
		return transactions, err
	}

	filteredTransactions := t.filter(query, transactions)

	return filteredTransactions, nil
}

func (t *TransactionPostgresService) filter(query map[string][]string, transactions []models.TransactionModel) []models.TransactionModel {
	//build out maps
	banks := make(map[string]bool)
	categories := make(map[string]bool)
	_, banksOk := query["bank"]
	if banksOk {
		t.filterHelper(strings.Split(query["bank"][0], ","), banks)
	}

	_, categoriesOk := query["category"]
	if categoriesOk {
		t.filterHelper(strings.Split(query["category"][0], ","), categories)
	}

	if len(banks) == 0 && len(categories) == 0 {
		return transactions
	}

	var filteredTransactions []models.TransactionModel

	for _, transaction := range transactions {
		validBank := banks[string(transaction.Bank)]

		validCategory := categories[string(transaction.Category)]
		if len(banks) > 0 && len(categories) == 0 {
			if validBank {
				filteredTransactions = append(filteredTransactions, transaction)
			}
			continue
		}

		if len(banks) == 0 && len(categories) > 0 {
			if validCategory {
				filteredTransactions = append(filteredTransactions, transaction)
			}
			continue
		}

		if len(banks) > 0 && len(categories) > 0 {
			if validBank && validCategory {
				filteredTransactions = append(filteredTransactions, transaction)
			}
		}

	}

	return filteredTransactions
}

func (t *TransactionPostgresService) filterHelper(from []string, to map[string]bool) {
	for i := range from {
		to[from[i]] = true
	}
}
