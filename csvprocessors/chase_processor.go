package csvprocessors

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

type Chase struct{}

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

func (c *Chase) ExtractCategories(records []*ChaseTransaction) {

}
