package services

import (
	"fmt"

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

func (t *TransactionPostgresService) generateID(bank categories.Bank, amount float64, category categories.Category, date, description string) string {
	return utils.EncodeToBase64(fmt.Sprintf("%s-%f-%s-%s-%s", bank, amount, category, date, description))
}
func (t *TransactionPostgresService) SearchTransactions(query map[string]interface{}) ([]models.TransactionModel, error) {

	return t.TransactionModel.SearchTransactions(query)
}
