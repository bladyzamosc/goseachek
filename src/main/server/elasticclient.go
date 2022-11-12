package server

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"goseachek/src/main/model"
	"log"
	"strconv"
	"sync"
)

type ElasticClient struct {
	es *elasticsearch.Client
}

const myindex string = "goseachek"

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

func (c ElasticClient) Index(index model.RequestIndex) {

	data, err := json.Marshal(index)
	if err != nil {
		log.Fatalf("Error marshaling document: %s", err)
	}

	req := esapi.IndexRequest{
		Index:      myindex,
		Body:       bytes.NewReader(data),
		Refresh:    "true",
		DocumentID: strconv.Itoa(1),
	}
	res, err := req.Do(context.Background(), c.es)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
}

func (c ElasticClient) Search(value string) string {

	var (
		r  map[string]interface{}
		wg sync.WaitGroup
	)

	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	res, err := c.es.Search(
		c.es.Search.WithContext(context.Background()),
		c.es.Search.WithIndex(myindex),
		c.es.Search.WithPretty(),
		c.es.Search.WithTrackTotalHits(true),
		c.es.Search.WithBody(&buf),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	wg.Wait()
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Print the response status, number of results, and request duration.
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r["took"].(float64)),
	)
	// Print the ID and document source for each hit.
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	}

	return "done"
}
