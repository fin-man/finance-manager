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

func (t *TransactionService) CreateTransaction(data []byte) error {

	return t.TransactionModel.CreateTransaction(data)
}

func (t *TransactionService) GetAllTransactions() *models.TransactionResponse {
	return t.TransactionModel.GetAllTransactions()
}

func (t *TransactionService) GetTransactionsInDateRange(from string, to string) *models.TransactionResponse {
	return t.TransactionModel.GetTransactionsInDateRange(from, to)
}
