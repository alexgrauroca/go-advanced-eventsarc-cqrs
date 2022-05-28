package eventsrepository

import (
	"context"
	"go-advanced-eventsarc-cqrs/models"
)

type EventStore interface {
	Close()

	// Feeds
	PublishCreatedFeed(ctx context.Context, feed *models.Feed) error
	SubscribeCreatedFeed(ctx context.Context) (<-chan CreatedFeedMessage, error)
	OnCreatedFeed(f func(CreatedFeedMessage)) error
}

var eventStore EventStore

func Close() {
	eventStore.Close()
}
