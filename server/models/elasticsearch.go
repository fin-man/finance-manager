package models

import (
	"finance-manager/elasticsearch"
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

}

func (e *ElasticSearchModel) CreateTransaction() {
	res, err := e.ElasticClient.Client.Index(
		"test",
		strings.NewReader(`{"transaction_date":"2020-03-29","amount":-16.23,"description":"AMZN Mktp US*IK34W8LC3","bank":"chase","account_id":"3341","category":"shopping"}`),
		e.ElasticClient.Client.Index.WithDocumentID("1"),
		e.ElasticClient.Client.Index.WithRefresh("true"),
	)
	if err != nil {
		log.Println("error when indexing : ", err)
	}

	defer res.Body.Close()

	log.Println(res)
}
