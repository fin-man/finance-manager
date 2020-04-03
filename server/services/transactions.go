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

func (t *TransactionService) GetAllTransactions() {
	t.TransactionModel.GetAllTransactions()
}
