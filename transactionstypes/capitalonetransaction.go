package transactionstypes

type CapitalOneTransaction struct {
	TransactionDate string `csv:"Transaction Date"`
	PostedDate string `csv:"Posted Date"`
	CardNo string `csv:"Card No."`
	Description  string `csv:"Description"`
	Category string `csv:"Category"`
	Debit string `csv:"Debit"`
	Credit string `csv:"Credit"`
}


func (c *CapitalOneTransaction) toString() string{
	return ""
}

