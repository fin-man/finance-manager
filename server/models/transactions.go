package models

import (
	"fmt"
	"time"

	"github.com/fin-man/finance-manager/categories"
	"github.com/fin-man/finance-manager/server/utils"
	"github.com/jinzhu/gorm"
)

type TransactionModel struct {
	gorm.Model
	TransactionID   string              `json:"transactions_id" gorm:"unique"`
	TransactionDate time.Time           `csv:"transaction_date" json:"transaction_date"`
	Amount          float64             `csv:"amount"  json:"amount"`
	Description     string              `csv:"description"  json:"description"`
	Bank            categories.Bank     `csv:"bank"  json:"bank"`
	AccountID       string              `csv:"account_id"  json:"account_id"`
	Category        categories.Category `csv:"category"  json:"category"`
	AccountType     string              `csv:"account_type" json:"account_type`
	Hash            string              `csv:"-" json:"hash"`
}

func NewTransactionsModel() *TransactionModel {
	return &TransactionModel{}
}

func (t TransactionModel) TableName() string {
	return "transactions"
}

func (e *TransactionModel) GetAllTransactions() ([]TransactionModel, error) {
	var transactions []TransactionModel
	if err := DB.Find(&transactions).Error; err != nil {
		return transactions, err
	}

	return transactions, err
}

/*
- /transactions
- /transactions?bank="bank_name"
- /transactions?category="category_name"
- /transactions?date="datetime"
- /transactions?start_time="start_datetime"&end_time="end_datetime"
-
*/
func (e *TransactionModel) SearchTransactions(query map[string][]string, startTime time.Time, endTime time.Time) ([]TransactionModel, error) {
	// map[string]interface{}{"name": []string{"jinzhu", "jinzhu 2"}}
	var transactions []TransactionModel
	// SELECT * FROM users WHERE name IN ('jinzhu','jinzhu 2');

	// fmt.Println("Bank ", query["bank"])
	// fmt.Println("Category ", query["category"])
	// if err := DB.Where("bank IN ?", []string{"bankOne", "bankTwo"}).Find(&transactions).Error; err != nil {
	// 	return transactions, err
	// }

	if err := DB.Where("transaction_date < ? AND transaction_date > ?", startTime, endTime).Find(&transactions).Error; err != nil {
		return transactions, err
	}

	return transactions, nil

}

func (e *TransactionModel) GetTransactionsInDateRange(search, from, to string) {

}

func (e *TransactionModel) CreateTransaction(transaction *TransactionModel) (*TransactionModel, error) {

	transactionID := e.generateTransactionID(transaction)
	transaction.TransactionID = transactionID
	if err := DB.Create(&transaction).Error; err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (e *TransactionModel) String() string {
	return fmt.Sprintf("%s-%f-%s-%s-%s-%s-%s", e.TransactionDate, e.Amount, e.Description, e.Bank, e.AccountID, e.Category, e.AccountType)
}
func (e *TransactionModel) generateTransactionID(transaction *TransactionModel) string {
	return utils.Sha1Hash(transaction.String())
}

// func (e *TransactionModel) generateTransactionQuery(query map[string][]string) string {
// 	for key , value := range query {

// 	}
// }
