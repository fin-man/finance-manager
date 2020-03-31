package csvprocessors

import (
	"finance-manager/transactionstypes"
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

type Chase struct{}

func NewChaseClient() *Chase {
	return &Chase{}
}

func (c *Chase) PrintRecords(records []*transactionstypes.ChaseTransaction, printCategories bool) {
	if printCategories {
		categories := make(map[string]bool)

		for _, record := range records {
			if record.Category != "" {
				categories[record.Category] = true
			}
		}

		for cat, _ := range categories {
			fmt.Println(cat)
		}
		return // skip this
	}

	for _, record := range records {
		fmt.Println(record)
	}
}

func (c *Chase) Unmarshal(file *os.File, records *[]*transactionstypes.ChaseTransaction) error {
	return gocsv.UnmarshalFile(file, records)
}
