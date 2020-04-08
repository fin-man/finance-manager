package categories

type NormalizedTransaction struct {
	TransactionDate string   `csv:"transaction_date" json:"transaction_date"`
	Amount          float64  `csv:"amount"  json:"amount"`
	Description     string   `csv:"description"  json:"description"`
	Bank            Bank     `csv:"bank"  json:"bank"`
	AccountID       string   `csv:"account_id"  json:"account_id"`
	Category        Category `csv:"category"  json:"category"`
	AccountType     string   `csv:"account_type" json:"account_type`
}

func (n *NormalizedTransaction) MakeAmountPositive() {
	if n.Amount < 0 {
		n.Amount *= -1
	}
}
