package events

import (
	"context"
	"go-advanced-eventsarc-cqrs/eventsrepository"
	"go-advanced-eventsarc-cqrs/models"

	"github.com/nats-io/nats.go"
)

func (n *NatsEventStore) PublishCreatedFeed(ctx context.Context, feed *models.Feed) error {
	msg := eventsrepository.CreatedFeedMessage{
		Id:          feed.Id,
		Title:       feed.Title,
		Description: feed.Description,
		CreatedAt:   feed.CreatedAt,
	}

	data, err := n.encodeMessage(msg)

	if err != nil {
		return err
	}

	return n.conn.Publish(msg.Type(), data)
}

func (n *NatsEventStore) OnCreatedFeed(f func(eventsrepository.CreatedFeedMessage)) (err error) {
	msg := eventsrepository.CreatedFeedMessage{}
	n.feedCreatedSub, err = n.conn.Subscribe(msg.Type(), func(m *nats.Msg) {
		n.decodeMessage(m.Data, &msg)
		f(msg)
	})

	return
}

func (n *NatsEventStore) SubscribeCreatedFeed(ctx context.Context) (<-chan eventsrepository.CreatedFeedMessage, error) {
	var err error

	m := eventsrepository.CreatedFeedMessage{}
	n.feedCreatedChan = make(chan eventsrepository.CreatedFeedMessage, 64)
	ch := make(chan *nats.Msg, 64)

	n.feedCreatedSub, err = n.conn.ChanSubscribe(m.Type(), ch)

	if err != nil {
		return nil, err
	}

	go func() {
		for {
			select {
			case msg := <-ch:
				n.decodeMessage(msg.Data, &m)
				n.feedCreatedChan <- m
			}
		}
	}()

	return (<-chan eventsrepository.CreatedFeedMessage)(n.feedCreatedChan), nil
}
