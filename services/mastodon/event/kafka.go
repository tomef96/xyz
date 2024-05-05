package event

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/tomef96/coop/mastodon/config"
)

type Event struct {
	Payload any    `json:"payload"`
	Schema  string `json:"schema"`
	Version int    `json:"version"`
}

type KafkaPublisher struct {
	kafka.Writer
}

func NewKafkaPublisher(topic string) *KafkaPublisher {
	CreateTopic(kafka.TopicConfig{
		Topic:             topic,
		NumPartitions:     3,
		ReplicationFactor: 1,
	})

	return &KafkaPublisher{kafka.Writer{
		Addr:         kafka.TCP(config.KAFKA_BROKER_URL),
		Topic:        topic,
		BatchTimeout: 30 * time.Millisecond,
	}}
}

func (p *KafkaPublisher) Publish(ctx context.Context, event Event) error {
	eventBytes, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = p.WriteMessages(ctx, kafka.Message{Value: eventBytes})
	if err != nil {
		log.Fatalf("failed to write messages: %v", err)
	}

	return nil
}

func CreateTopic(topicConfig kafka.TopicConfig) {
	var conn *kafka.Conn
	for i := 0; i < 10; i++ {
		innerConn, err := kafka.Dial("tcp", config.KAFKA_BROKER_URL)
		if err != nil {
			log.Printf("failed to dial %s, retrying in 1 second", config.KAFKA_BROKER_URL)
			time.Sleep(time.Second)
			continue
		}
		conn = innerConn
		defer conn.Close()
	}

	if conn == nil {
		log.Fatalf("failed to dial kafka broker")
	}

	controller, err := conn.Controller()
	if err != nil {
		log.Fatalf("failed to resolve kafka controller: %v", err)
	}

	controllerConn, err := kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		log.Fatalf("failed to dial kafka controller: %v", err)
	}
	defer controllerConn.Close()

	err = controllerConn.CreateTopics(topicConfig)
	if err != nil {
		log.Fatalf("failed to create kafka topics: %v", err)
	}
}
