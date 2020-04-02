package elasticsearch 

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)
type ElasticSearchClient struct {
	Client *elasticsearch.Client
}

func NewElasticSearchClient() *ElasticSearchClient{

	es , err := elasticsearch.NewDefaultClient()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("ElasticSearch client initialized")

	log.Println("Checking connection .. ")
	res , err := es.Info()
	if err != nil{
		log.Fatalf("Unable to establish connection with ElasticSearch API : %v \n", err)
	}

	log.Printf("Conection established successfully .. : %s\n",res.String())

	
	return &ElasticSearchClient{
		Client: es,
	}
}


