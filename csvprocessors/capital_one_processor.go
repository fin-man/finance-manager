package csvprocessors

import (
	"finance-manager/transactionstypes"
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

type CapitalOne struct{}

func NewCapitalOneClient() *CapitalOne {
	return &CapitalOne{}
}

func (c *CapitalOne) PrintRecords(records []*transactionstypes.CapitalOneTransaction, printCategories bool) {

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

func (c *CapitalOne) Unmarshal(file *os.File, records *[]*transactionstypes.CapitalOneTransaction) error {
	return gocsv.UnmarshalFile(file, records)
}
