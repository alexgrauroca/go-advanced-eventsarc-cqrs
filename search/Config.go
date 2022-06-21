package search

type Config struct {
	ElasticsearchAddress string `envconfig:"ELASTICSEARCH_ADDRESS"`
}
