package main

import (
	"log"

	"github.com/tomef96/coop/mastodon/domain"
	"github.com/tomef96/coop/mastodon/event"
	"github.com/tomef96/coop/mastodon/storage"
	"github.com/tomef96/coop/mastodon/transport"
)

func main() {
	postsStorage := storage.NewPostsStorage()
	defer postsStorage.Close()
	postsPublisher := event.NewPostsPublisher()
	defer postsPublisher.Close()
	postsService := domain.NewPostsService(postsStorage, postsPublisher)

	router := transport.Router(postsService)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
