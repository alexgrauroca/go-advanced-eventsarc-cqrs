package main

import (
	"go-advanced-eventsarc-cqrs/database"
	"go-advanced-eventsarc-cqrs/events"
	"go-advanced-eventsarc-cqrs/eventsrepository"
	"go-advanced-eventsarc-cqrs/repository"
	"log"
	"net/http"
)

func main() {
	repo, err := database.NewPostgresRepository("")

	if err != nil {
		log.Fatal(err)
	}

	repository.SetRepository(repo)
	defer repository.Close()

	ev, err := events.NewNats("")

	if err != nil {
		log.Fatal(err)
	}

	eventsrepository.SetEventStore(ev)
	defer eventsrepository.Close()

	router := newRouter()

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
