package searchrepository

import (
	"context"
	"go-advanced-eventsarc-cqrs/models"
)

func IndexFeed(ctx context.Context, feed models.Feed) error {
	return repo.IndexFeed(ctx, feed)
}

func SearchFeed(ctx context.Context, query string) ([]models.Feed, error) {
	return repo.SearchFeed(ctx, query)
}
