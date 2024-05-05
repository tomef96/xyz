package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/tomef96/coop-test/log/config"
	"github.com/tomef96/coop-test/log/event"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	postsConsumer := event.NewConsumer(config.KAFKA_TOPIC_POSTS)
	defer postsConsumer.Close()

	go postsConsumer.Consume(ctx)

	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel
	cancel()
}
