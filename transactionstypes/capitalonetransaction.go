package transactionstypes

import "fmt"

type CapitalOneTransaction struct {
	TransactionDate string `csv:"Transaction Date"`
	PostedDate      string `csv:"Posted Date"`
	CardNo          string `csv:"Card No."`
	Description     string `csv:"Description"`
	Category        string `csv:"Category"`
	Debit           string `csv:"Debit"`
	Credit          string `csv:"Credit"`
}

func (c *CapitalOneTransaction) String() string {
	return fmt.Sprintf("TransactionDate : %s \n PostedDate : %s \n CardNo : %s \n Description : %s\n Category : % s \n Debit : %s \n Credit : %s \n", c.TransactionDate, c.PostedDate, c.CardNo, c.Description, c.Category, c.Debit, c.Credit)
}
