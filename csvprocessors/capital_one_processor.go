package csvprocessors

import (
	"finance-manager/categories"
	"finance-manager/transactionstypes"
	"finance-manager/utils"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gocarina/gocsv"
)

type CapitalOne struct{}

func NewCapitalOneClient() *CapitalOne {
	return &CapitalOne{}
}

func (c *CapitalOne) ProcessCSV(records []*transactionstypes.CapitalOneTransaction) []*categories.NormalizedTransaction {
	var normalizedRecords []*categories.NormalizedTransaction

	for _, record := range records {

		formatedTime, err := c.ConverTime(record.TransactionDate)
		if err != nil {
			log.Printf("Unable to convert time in CapitalOne for record : %s \n", record.String())
			//TO DO : ALERT EXTERNALLY
			continue //skip the record
		}

		amountToFloat, err := strconv.ParseFloat(record.Debit, 64)

		if err != nil {
			log.Printf("Unable to convert the amount in CapitalOne for record : %s \n", record.String())
			//TO DO : ALERT EXTERNALLY
			continue //skip the record
		}

		pickedCategory := categories.CapitalOneTransactionTypes[record.Category]

		normalizedRecord := categories.NormalizedTransaction{
			TransactionDate: formatedTime,
			Amount:          amountToFloat,
			Description:     record.Description,
			Bank:            categories.CapitalOne,
			AccountID:       record.CardNo,
			Category:        pickedCategory,
		}

		normalizedRecords = append(normalizedRecords, &normalizedRecord)
	}

	return normalizedRecords
}

func (c *CapitalOne) ConverTime(csvTimeStamp string) (string, error) {
	t, err := time.Parse(utils.TimeLayout, csvTimeStamp)

	if err != nil {
		return "", err
	}
	return t.Format(utils.TimeLayout), err
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
