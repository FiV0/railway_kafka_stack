package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Kafka producer example, following Confluent's documentation at https://docs.confluent.io/kafka-clients/go/current/overview.html#ak-producer
func kafkaProducer(topic string, config *kafka.ConfigMap) {
	p, err := kafka.NewProducer(config)
	if err != nil {
		log.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	counter := 0
	for {
		<-time.After(1 * time.Second) // Produce a message every second
		message := fmt.Sprintf("Example message #%d from producer", counter)
		jsonmsg, _ := json.Marshal(message)
		err = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          jsonmsg,
		},
			nil, // delivery channel
		)
		if err != nil {
			log.Printf("failed to produce message: %s", err)
			break
		}
		fmt.Printf("sent message: '%s'\n", message)
		counter += 1
	}
}
