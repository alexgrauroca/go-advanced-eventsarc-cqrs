package main

import (
	"go-advanced-eventsarc-cqrs/events"
	"go-advanced-eventsarc-cqrs/eventsrepository"
	"log"
	"net/http"
)

func main() {
	n, err := events.NewNats("")

	if err != nil {
		log.Fatal(err)
	}

	hub := NewHub()

	err = n.OnCreatedFeed(func(m eventsrepository.CreatedFeedMessage) {
		hub.Broadcast(newCreatedFeedMessage(m.Id, m.Title, m.Description, m.CreatedAt), nil)
	})

	if err != nil {
		log.Fatal(err)
	}

	eventsrepository.SetEventStore(n)
	defer eventsrepository.Close()

	go hub.Run()

	http.HandleFunc("/ws", hub.HandleWebSocket)
	err = http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
