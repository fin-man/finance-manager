package services

import (
	"github.com/fin-man/finance-manager/server/models"
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
	return t.TransactionModel.CreateTransaction(transaction)
}
