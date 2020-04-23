package services

type CSVService struct {
}

func NewCSVService() *CSVService {
	return &CSVService{}
}

func (c *CSVService) GetFields() []string {

	return []string{
		"transaction_date",
		"account_id",
		"account_type",
		"bank",
		"description",
		"category",
		"amount",
	}
}
