package categories

type NormalizedTransaction struct{
	TransactionDate string `csv:"transaction_date" json:"transaction_date"`
	Amount float64 `csv:"amount"  json:"amount"`
	Description string `csv:"description"  json:"description"`
	Bank Bank  `csv:"bank"  json:"bank"`
	AccountID string `csv:"account"  json:"account_id"`
	Category Category `csv:"category"  json:"category"`
}