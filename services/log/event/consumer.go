package event

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/tomef96/coop-test/log/config"
)

func Consume(topic string) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{config.KAFKA_BROKER_URL},
		Topic:   topic,
		GroupID: fmt.Sprintf("%v-consumers", topic),
	})
	defer reader.Close()

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("error while reading message: %v\n", err)
			continue
		}

		fmt.Printf(
			"message at topic/partition/offset %v/%v/%v: %s = %s\n",
			m.Topic,
			m.Partition,
			m.Offset,
			string(m.Key),
			string(m.Value),
		)
	}
}
