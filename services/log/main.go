package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/tomef96/coop-test/log/config"
	"github.com/tomef96/coop-test/log/event"
)

func main() {
	go event.Consume(config.KAFKA_TOPIC_POSTS)

	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel
}
