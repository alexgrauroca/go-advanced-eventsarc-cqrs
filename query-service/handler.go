package main

import (
	"context"
	"go-advanced-eventsarc-cqrs/eventsrepository"
	"go-advanced-eventsarc-cqrs/models"
	"go-advanced-eventsarc-cqrs/searchrepository"
	"log"
)

func onCreatedFeed(m eventsrepository.CreatedFeedMessage) {
	feed := models.Feed{
		Id:          m.Id,
		Title:       m.Title,
		Description: m.Description,
		CreatedAt:   m.CreatedAt,
	}

	if err := searchrepository.IndexFeed(context.Background(), feed); err != nil {
		log.Printf("Failed to index feed: %v", err)
	}
}
