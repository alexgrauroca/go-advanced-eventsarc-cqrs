package search

import (
	"bytes"
	"context"
	"encoding/json"
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
