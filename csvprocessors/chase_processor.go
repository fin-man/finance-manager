package csvprocessors

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

type Chase struct{}

type ChaseTransaction struct {
	Transaction string `csv:"Transaction Date"`
	PostDate    string `csv:"Post Date"`
	Description string `csv:"Description"`
	Category    string `csv:"Category"`
	Type        string `csv:"Type"`
	Amount      string `csv:"Amount"`
}

func NewChaseClient() *Chase {
	return &Chase{}
}

func (c *Chase) PrintRecords(records []*ChaseTransaction) {
	for _, record := range records {

		fmt.Println(record)
	}
}

func (c *Chase) Unmarshal(file *os.File, records *[]*ChaseTransaction) error {
	return gocsv.UnmarshalFile(file, records)
}
