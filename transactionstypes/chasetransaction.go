package transactionstypes

type ChaseTransaction struct {
	Transaction string `csv:"Transaction Date"`
	PostDate    string `csv:"Post Date"`
	Description string `csv:"Description"`
	Category    string `csv:"Category"`
	Type        string `csv:"Type"`
	Amount      string `csv:"Amount"`
}

func (c *ChaseTransaction) ToString() string {

}
