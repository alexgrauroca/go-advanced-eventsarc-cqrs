package events

import "time"

type CreatedFeedMessage struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

func (m *CreatedFeedMessage) Type() string {
	return "created_feed"
}
