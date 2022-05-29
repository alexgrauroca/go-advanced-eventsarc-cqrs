package events

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"go-advanced-eventsarc-cqrs/eventsrepository"

	"github.com/kelseyhightower/envconfig"
	"github.com/nats-io/nats.go"
)

type NatsEventStore struct {
	conn            *nats.Conn
	feedCreatedSub  *nats.Subscription
	feedCreatedChan chan eventsrepository.CreatedFeedMessage
}

func NewNats(url string) (*NatsEventStore, error) {
	if len(url) == 0 {
		var err error
		url, err = loadServerUrl()

		if err != nil {
			return nil, err
		}
	}

	conn, err := nats.Connect(url)

	if err != nil {
		return nil, err
	}

	return &NatsEventStore{
		conn: conn,
	}, nil
}

func loadServerUrl() (string, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("nats://%s", cfg.NatsAddress), nil
}

func (n *NatsEventStore) Close() {
	if n.conn != nil {
		n.conn.Close()
	}

	if n.feedCreatedSub != nil {
		n.feedCreatedSub.Unsubscribe()
	}

	close(n.feedCreatedChan)
}

func (n *NatsEventStore) encodeMessage(m eventsrepository.Message) ([]byte, error) {
	b := bytes.Buffer{}
	err := gob.NewEncoder(&b).Encode(m)

	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func (n *NatsEventStore) decodeMessage(data []byte, m any) error {
	b := bytes.Buffer{}
	b.Write(data)

	return gob.NewDecoder(&b).Decode(m)
}
