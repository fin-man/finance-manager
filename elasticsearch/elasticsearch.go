package elasticsearch

import (
	"log"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
)

type ElasticSearchClient struct {
	Client *elasticsearch.Client
}

func NewElasticSearchClient() *ElasticSearchClient {

	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
		// ...
	}

	es, err := elasticsearch.NewClient(cfg)

	es.Ping.WithPretty()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("ElasticSearch client initialized")

	tries := 0
	for tries <= 10 {
		res, err := es.Ping()

		if res != nil {
			break
		}
		tries++
		log.Printf("Unable to establish connection with ElasticSearch Cluster. Error : %v ,  Retry count : %d\n", err, tries)
		time.Sleep(2 * time.Second)
	}

	res, err := es.Info()
	if err != nil {
		log.Fatalf("Still Unable to establish connection with ElasticSearch API : %v \n", err)
	}

	log.Printf("Conection established successfully .. : %s\n", res.String())

	return &ElasticSearchClient{
		Client: es,
	}
}
