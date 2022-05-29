package searchrepository

import (
	"context"
	"go-advanced-eventsarc-cqrs/models"
)

type SearchRepository interface {
	Close()
	IndexFeed(ctx context.Context, feed models.Feed) error
	SearchFeed(ctx context.Context, query string) ([]models.Feed, error)
}

var repo SearchRepository

func SetSearchRepository(r SearchRepository) {
	repo = r
}

func Close() {
	repo.Close()
}
