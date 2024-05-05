package config

import (
	"log"
	"os"
)

func requireString(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("the environment variable '%s' must be specified", key)
	}

	return value
}

var (
	KAFKA_BROKER_URL  = requireString("KAFKA_BROKER_URL")
	KAFKA_TOPIC_POSTS = requireString("KAFKA_TOPIC_POSTS")
	MONGO_URL         = requireString("MONGO_URL")
	MONGO_USER        = requireString("MONGO_USER")
	MONGO_PASS        = requireString("MONGO_PASS")
)
