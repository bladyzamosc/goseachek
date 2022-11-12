package server

import (
	"github.com/elastic/go-elasticsearch/v8"
	"log"
)

type ElasticClient struct {
	es *elasticsearch.Client
}

func (ElasticClient) NewElasticClient() ElasticClient {
	es, er := elasticsearch.NewClient(ElasticConfig())
	if er != nil {
		log.Fatalf("Error creating the client: %s", er)
	}
	res, er := es.Info()
	if er != nil {
		log.Fatalf("Error getting response: %s", er)
	}
	defer res.Body.Close()
	log.Println(res)
	ec := ElasticClient{}
	ec.es = es
	return ec
}

func ElasticConfig() elasticsearch.Config {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
		Username: "elastic",
		Password: "admin123",
	}
	return cfg
}
