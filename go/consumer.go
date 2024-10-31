package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Kafka consumer example, following Confluent's documentation at https://docs.confluent.io/kafka-clients/go/current/overview.html#ak-consumer
func kafkaConsumer(topic string, config *kafka.ConfigMap) {
	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		log.Fatalf("Failed to create consumer: %s\n", err)
	}

	err = consumer.Subscribe(topic, nil)
	if err != nil {
		log.Fatalf("Failed to create consumer: %s\n", err)
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		content := ""
		if err == nil {
			json.Unmarshal(msg.Value, &content)
			fmt.Printf("consumed message from topic: %s\n", content)
		} else {
			log.Printf("consumer failed with kafka error: %s", err)
			break
		}
	}

	consumer.Close()
}
