package main

import (
	"sync"

	"github.com/elastic/go-elasticsearch/v7"
)

var (
	es *elasticsearch.Client
	mu sync.Mutex
)

func GetEsClient() (*elasticsearch.Client, error) {
	if es != nil {
		return es, nil
	}

	mu.Lock()
	defer mu.Unlock()
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}
	es7, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	es = es7

	return es, nil
}
