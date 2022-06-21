package main

import (
	"go-advanced-eventsarc-cqrs/database"
	"go-advanced-eventsarc-cqrs/events"
	"go-advanced-eventsarc-cqrs/eventsrepository"
	"go-advanced-eventsarc-cqrs/repository"
	"go-advanced-eventsarc-cqrs/search"
	"go-advanced-eventsarc-cqrs/searchrepository"
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

	es, err := search.NewElastic("")

	if err != nil {
		log.Fatal(err)
	}

	log.Println(es)

	searchrepository.SetSearchRepository(es)
	defer searchrepository.Close()

	n, err := events.NewNats("")

	if err != nil {
		log.Fatal(err)
	}

	err = n.OnCreatedFeed(onCreatedFeed)

	if err != nil {
		log.Fatal(err)
	}

	eventsrepository.SetEventStore(n)
	defer eventsrepository.Close()

	router := newRouter()

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
