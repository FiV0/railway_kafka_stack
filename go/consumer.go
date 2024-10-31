package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Kafka consumer example, following Confluent's documentation at https://docs.confluent.io/kafka-clients/go/current/overview.html#ak-consumer
func kafkaConsumer(topic string, config *kafka.ConfigMap) {
	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		log.Printf("Failed to create consumer: %s\n", err)
		os.Exit(1)
	}

	err = consumer.Subscribe(topic, nil)
	if err != nil {
		log.Printf("Failed to create consumer: %s\n", err)
		os.Exit(1)
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		content := ""
		if err == nil {
			json.Unmarshal(msg.Value, &content)
			log.Printf("consumed message from topic: %s", content)
		} else {
			log.Printf("consumer failed with kafka error: %s", err)
			break
		}
	}

	consumer.Close()
}
