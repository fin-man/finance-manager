package models

import (
	"github.com/jinzhu/gorm"
)

type TransactionModel struct {
}

type TransactionResponse struct {
	gorm.
}

func NewTransactionsModel() *TransactionModel {
	return &TransactionModel{}
}

func (e *TransactionModel) GetAllTransactions() {

}

func (e *TransactionModel) SearchTransaction(search, from, to string) {

}

func (e *TransactionModel) GetTransactionsInDateRange(search, from, to string) {

}

func (e *TransactionModel) CreateTransaction(data []byte, id string) {

}
