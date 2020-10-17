package services

/*
type TransactionService struct {
	TransactionModel *models.ElasticSearchModel
}

type GraphFullResponse struct {
	AllTransactions *models.TransactionResponse `json:"all_transactions"`
	Graph           GraphResponse               `json:"graph_response"`
}

type GraphResponse struct {
	AllGraph            [][]float64                                                `json:"all_graph_data"`
	CategoryMap         map[categories.Category][][]float64                        `json"category_graph_data"`
	CategoryMapDetailed map[categories.Category][]categories.NormalizedTransaction `json:"category_map_detailed"`
}

func NewTransactionService() *TransactionService {
	transactionModel := models.NewElasticSearchModel()
	return &TransactionService{
		TransactionModel: transactionModel,
	}
}

func (t *TransactionService) CreateTransaction(data []byte, id string) error {

	return t.TransactionModel.CreateTransaction(data, id)
}

func (t *TransactionService) SearchTransaction(search, from, to string) (*models.TransactionResponse, error) {
	return t.TransactionModel.SearchTransaction(search, from, to)
}
func (t *TransactionService) GetAllTransactions() (*models.TransactionResponse, error) {
	return t.TransactionModel.GetAllTransactions()
}

func (t *TransactionService) GetAllTransactionsGraph(from string, to string) *GraphFullResponse {
	transactions, err := t.TransactionModel.GetTransactionsInDateRange(from, to)
	if err != nil {
		//abort
	}

	var graphFullResponse GraphFullResponse
	graphFullResponse.AllTransactions = transactions

	var allGraphData [][]float64

	timeStampTotals := make(map[float64]float64)
	categoryTimeStamp := make(map[categories.Category]map[float64]float64)
	categoryDetailed := make(map[categories.Category][]categories.NormalizedTransaction)
	for _, t := range transactions.Hits.Hits {

		_, okCat := categoryTimeStamp[t.Source.Category]
		if !okCat {
			categoryTimeStamp[t.Source.Category] = make(map[float64]float64)
		}

		categoryDetailed[t.Source.Category] = append(categoryDetailed[t.Source.Category], t.Source)

		unixMili, err := utils.ConvertTimeToUnixMillis(t.Source.TransactionDate)

		if err != nil {
			log.Println(err)
			continue
		}

		_, ok := timeStampTotals[float64(unixMili)]

		if t.Source.Amount < 0 {
			t.Source.Amount = t.Source.Amount * -1
		}

		if !ok {
			//first time
			timeStampTotals[float64(unixMili)] = t.Source.Amount
			categoryTimeStamp[t.Source.Category][float64(unixMili)] = t.Source.Amount
			continue
		}

		timeStampTotals[float64(unixMili)] += t.Source.Amount
		categoryTimeStamp[t.Source.Category][float64(unixMili)] += t.Source.Amount
	}

	for date, amount := range timeStampTotals {
		record := make([]float64, 2)
		record[0] = date
		record[1] = amount

		allGraphData = append(allGraphData, record)
	}

	var graphResponse GraphResponse

	sort.Slice(allGraphData, func(i, j int) bool {
		return allGraphData[i][0] < allGraphData[j][0]
	})

	graphResponse.AllGraph = allGraphData

	categoryMap := make(map[categories.Category][][]float64)
	for category, dailyValues := range categoryTimeStamp {
		var tempDateValues [][]float64
		//categoryMap[category] = tempDateValues

		for date, amount := range dailyValues {
			record := make([]float64, 2)
			record[0] = date
			record[1] = amount

			tempDateValues = append(tempDateValues, record)
		}

		sort.Slice(tempDateValues, func(i, j int) bool {
			return tempDateValues[i][0] < tempDateValues[j][0]
		})

		categoryMap[category] = tempDateValues
	}

	for _, detailedData := range categoryDetailed {
		sort.Slice(detailedData, func(i, j int) bool {
			return detailedData[i].TransactionDate < detailedData[j].TransactionDate
		})
	}
	graphResponse.CategoryMap = categoryMap
	graphResponse.CategoryMapDetailed = categoryDetailed

	graphFullResponse.Graph = graphResponse

	return &graphFullResponse
}

func (t *TransactionService) GetTransactionsInDateRange(from string, to string) (*models.TransactionResponse, error) {
	return t.TransactionModel.GetTransactionsInDateRange(from, to)
}
*/
