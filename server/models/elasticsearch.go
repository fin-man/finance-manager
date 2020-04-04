package models

import (
	"bytes"
	"context"
	"encoding/json"
	"finance-manager/categories"
	"finance-manager/elasticsearch"
	"fmt"
	"log"
	"strings"
)

type ElasticSearchModel struct {
	ElasticClient *elasticsearch.ElasticSearchClient
}

type TransactionResponse struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []Hit   `json:"hits"`
	} `json:"hits"`
}

type Hit struct {
	Index  string                           `json:"_index"`
	Type   string                           `json:"_type"`
	ID     string                           `json:"_id"`
	Score  float64                          `json:"_score"`
	Source categories.NormalizedTransaction `json:"_source,omitempty"`
}

func NewElasticSearchModel() *ElasticSearchModel {

	elasticClient := elasticsearch.NewElasticSearchClient()
	return &ElasticSearchModel{
		ElasticClient: elasticClient,
	}

}

func (e *ElasticSearchModel) GetAllTransactions() (*TransactionResponse, error) {
	var buf bytes.Buffer
	var tr TransactionResponse

	query := map[string]interface{}{
		"from": 0,
		"size": 10000,
		"sort": []map[string]map[string]string{{"transaction_date": {"order": "desc"}}},
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return &tr, err
	}

	// Perform the search request.
	res, err := e.ElasticClient.Client.Search(
		e.ElasticClient.Client.Search.WithContext(context.Background()),
		e.ElasticClient.Client.Search.WithIndex("transactions"),
		e.ElasticClient.Client.Search.WithBody(&buf),
		e.ElasticClient.Client.Search.WithTrackTotalHits(true),
		e.ElasticClient.Client.Search.WithPretty(),
	)

	if err != nil {
		return &tr, err
	}

	readBuf := new(bytes.Buffer)

	readBuf.ReadFrom(res.Body)

	err = json.Unmarshal(readBuf.Bytes(), &tr)

	if err != nil {
		return &tr, err
	}

	return &tr, err
}

func (e *ElasticSearchModel) GetTransactionsInDateRange(from string, to string) (*TransactionResponse, error) {
	var buf bytes.Buffer
	var tr TransactionResponse

	query := map[string]interface{}{
		"from": 0,
		"size": 10000,
		"query": map[string]interface{}{
			"range": map[string]map[string]string{"transaction_date": {"gte": from, "lte": to}},
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return &tr, err
	}

	fmt.Println(buf.String())
	// Perform the search request.
	res, err := e.ElasticClient.Client.Search(
		e.ElasticClient.Client.Search.WithContext(context.Background()),
		e.ElasticClient.Client.Search.WithIndex("transactions"),
		e.ElasticClient.Client.Search.WithBody(&buf),
		e.ElasticClient.Client.Search.WithTrackTotalHits(true),
		e.ElasticClient.Client.Search.WithPretty(),
	)

	if err != nil {
		return &tr, err

	}

	readBuf := new(bytes.Buffer)

	readBuf.ReadFrom(res.Body)

	err = json.Unmarshal(readBuf.Bytes(), &tr)

	if err != nil {
		return &tr, err
	}

	return &tr, err
}

func (e *ElasticSearchModel) CreateTransaction(data []byte) error {
	res, err := e.ElasticClient.Client.Index(
		"transactions",
		strings.NewReader(string(data)),
		e.ElasticClient.Client.Index.WithRefresh("true"),
	)
	if err != nil {
		log.Println("error when indexing : ", err)
		return err
	}

	defer res.Body.Close()

	log.Println(res)
	return err
}
