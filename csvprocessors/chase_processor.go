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

type Chase struct {
}

var (
	ChaseTimeStampLayout string = "01/02/2006"
)

func NewChaseClient() *Chase {
	return &Chase{}
}

func (c *Chase) ProcessCSV(records []*transactionstypes.ChaseTransaction) []*categories.NormalizedTransaction {
	var normalizedRecords []*categories.NormalizedTransaction

	for _, record := range records {
		formatedTime, err := c.ConverTime(record.TransactionDate)

		if err != nil {
			log.Printf("Unable to convert time in CapitalOne for record : %s \n", record.String())
			//TO DO : ALERT EXTERNALLY
			continue //skip the record
		}

		amountToFloat, err := strconv.ParseFloat(record.Amount, 64)
		if err != nil {
			log.Printf("Unable to convert the amount in CapitalOne for record : %s \n", record.String())
			//TO DO : ALERT EXTERNALLY
			continue //skip the record
		}

		pickedCategory, ok := categories.ChaseTransactionTypes[record.Category]

		if !ok {
			log.Printf("UnRecognized category in Chase record : %s \n", record.String())
			log.Println("Category : ", record.Category)
			continue
		}
		normalizedRecord := categories.NormalizedTransaction{
			TransactionDate: formatedTime,
			Amount:          amountToFloat,
			Description:     record.Description,
			Bank:            categories.Chase,
			AccountID:       "3341",
			Category:        pickedCategory,
		}

		normalizedRecords = append(normalizedRecords, &normalizedRecord)

	}
	return normalizedRecords
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

func (c *Chase) ConverTime(csvTimeStamp string) (string, error) {
	t, err := time.Parse(ChaseTimeStampLayout, csvTimeStamp)

	if err != nil {
		return "", err
	}
	return t.Format(utils.TimeLayout), err
}

func (c *Chase) Unmarshal(file *os.File, records *[]*transactionstypes.ChaseTransaction) error {
	return gocsv.UnmarshalFile(file, records)
}
