package main

import (
	"context"
	"encoding/json"
	"go-advanced-eventsarc-cqrs/eventsrepository"
	"go-advanced-eventsarc-cqrs/models"
	"go-advanced-eventsarc-cqrs/repository"
	"go-advanced-eventsarc-cqrs/searchrepository"
	"log"
	"net/http"
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

func listFeedsHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	ctx := r.Context()

	feeds, err := repository.ListFeeds(ctx)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(feeds)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	ctx := r.Context()
	query := r.URL.Query().Get("q")

	if len(query) == 0 {
		http.Error(w, "query required", http.StatusBadRequest)
		return
	}

	feeds, err := searchrepository.SearchFeed(ctx, query)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(feeds)
}
