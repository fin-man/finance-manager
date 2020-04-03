package services

import "finance-manager/server/models"

type TransactionService struct {
	TransactionModel *models.ElasticSearchModel
}

func NewTransactionService() *TransactionService {
	transactionModel := models.NewElasticSearchModel()
	return &TransactionService{
		TransactionModel: transactionModel,
	}
}

func (t *TransactionService) CreateTransaction() {
	t.TransactionModel.CreateTransaction()
}
