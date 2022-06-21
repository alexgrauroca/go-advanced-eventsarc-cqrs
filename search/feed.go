package search

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"go-advanced-eventsarc-cqrs/models"
)

func (r *ElasticSearchRepository) IndexFeed(ctx context.Context, feed models.Feed) (err error) {
	body, _ := json.Marshal(feed)
	_, err = r.client.Index(
		"feeds",
		bytes.NewReader(body),
		r.client.Index.WithDocumentID(feed.Id),
		r.client.Index.WithContext(ctx),
		r.client.Index.WithRefresh("wait_for"),
	)
	return
}

func (r *ElasticSearchRepository) SearchFeed(ctx context.Context, query string) ([]models.Feed, error) {
	var buf bytes.Buffer
	searchQuery := map[string]any{
		"query": map[string]any{
			"multi_match": map[string]any{
				"query":            query,
				"fields":           []string{"title", "description"},
				"fuzziness":        3,
				"cutoff_frequency": 0.0001,
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(searchQuery); err != nil {
		return nil, err
	}

	res, err := r.client.Search(
		r.client.Search.WithContext(ctx),
		r.client.Search.WithIndex("feeds"),
		r.client.Search.WithBody(&buf),
		r.client.Search.WithTrackTotalHits(true),
	)

	if err != nil {
		return nil, err
	}

	var feeds []models.Feed

	defer func() {
		if err := res.Body.Close(); err != nil {
			feeds = nil
		}
	}()

	if res.IsError() {
		return nil, errors.New(res.String())
	}

	var eRes map[string]any

	if err := json.NewDecoder(res.Body).Decode(&eRes); err != nil {
		return nil, err
	}

	for _, hit := range eRes["hits"].(map[string]any)["hits"].([]any) {
		feed := models.Feed{}
		source := hit.(map[string]any)["_source"]
		marshal, err := json.Marshal(source)

		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(marshal, &feed); err == nil {
			feeds = append(feeds, feed)
		}
	}

	return feeds, nil
}
