package models

import (
	"bytes"
	"context"
	"encoding/json"
	"finance-manager/elasticsearch"
	"fmt"
	"log"
	"strings"
)

type ElasticSearchModel struct {
	ElasticClient *elasticsearch.ElasticSearchClient
}

func NewElasticSearchModel() *ElasticSearchModel {

	elasticClient := elasticsearch.NewElasticSearchClient()
	return &ElasticSearchModel{
		ElasticClient: elasticClient,
	}

}

func (e *ElasticSearchModel) GetAllTransactions() {

	// ctx := context.Background()
	// res, err := e.ElasticClient.Client.Search(
	// 	e.ElasticClient.Client.Search.WithContext(ctx),
	// 	e.ElasticClient.Client.withInd
	// )

	// var buf bytes.Buffer
	// query := map[string]interface{}{
	// 	"query": map[string]interface{}{
	// 		"match": map[string]interface{}{
	// 			"description": "AMZN",
	// 		},
	// 	},
	// }

	var buf bytes.Buffer
	query := map[string]interface{}{
		"from": 0,
		"size": 10000,
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
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
		log.Printf("ERROR : %v \n", err)
	}

	fmt.Printf("RESPONSE : %v \n ", res)
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
