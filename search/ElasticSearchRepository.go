package search

import (
	"fmt"

	elastic "github.com/elastic/go-elasticsearch/v7"
	"github.com/kelseyhightower/envconfig"
)

type ElasticSearchRepository struct {
	client *elastic.Client
}

func NewElastic(url string) (*ElasticSearchRepository, error) {
	if len(url) == 0 {
		var err error
		url, err = loadServerUrl()

		if err != nil {
			return nil, err
		}
	}

	client, err := elastic.NewClient(elastic.Config{
		Addresses: []string{url},
	})

	if err != nil {
		return nil, err
	}

	return &ElasticSearchRepository{client}, nil
}

func loadServerUrl() (string, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("http://%s", cfg.ElasticsearchAddress), nil
}

func (r *ElasticSearchRepository) Close() {}
