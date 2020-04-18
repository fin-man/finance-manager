package csvprocessors

import (
	"finance-manager/categories"
	"os"

	"github.com/gocarina/gocsv"
)

type CSVProcessor struct{}

var (
	TimeStampLayout string = "2006-01-02"
)

func NewCSVprocessor() *CSVProcessor {
	return &CSVProcessor{}
}

func (c *CSVProcessor) Unmarshal(file *os.File, records *[]*categories.NormalizedTransaction) error {
	gocsv.FailIfUnmatchedStructTags = true
	return gocsv.UnmarshalFile(file, records) // gets this unmarshaleld
}
