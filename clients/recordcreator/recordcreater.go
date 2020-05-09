package recordcreator

import (
	"encoding/json"

	"github.com/fin-man/finance-manager/requests"

	"github.com/fin-man/finance-manager/categories"
)

type RecordCreator struct {
	Requests *requests.Requests
}

func NewRecordCreator() *RecordCreator {

	return &RecordCreator{
		Requests: requests.NewRequestsClient(),
	}

}

func (r *RecordCreator) CreateNewRecord(record *categories.NormalizedTransaction) error {

	url := "http://localhost:8080/transactions"

	body, err := json.Marshal(record)

	if err != nil {
		return err
	}

	err = r.Requests.NewPOSTRquest(url, body)

	if err != nil {
		return err
	}

	return nil
}
