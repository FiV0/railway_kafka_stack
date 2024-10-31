package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Kafka producer example, following Confluent's documentation at https://docs.confluent.io/kafka-clients/go/current/overview.html#ak-producer
func kafkaProducer(topic string, config *kafka.ConfigMap) {
	p, err := kafka.NewProducer(config)
	if err != nil {
		log.Fatalf("Failed to create producer: %s\n", err)
	}

	counter := 0
	for {
		<-time.After(1 * time.Second) // Produce a message every second
		message := fmt.Sprintf("Example message #%d from producer", counter)
		jsonmsg, _ := json.Marshal(message)

		deliveries := make(chan kafka.Event)
		err = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          jsonmsg,
		}, deliveries,
		)
		if err != nil {
			log.Fatalf("failed to produce message: %s", err)
		}

		select {
		case event := <-deliveries:
			fmt.Printf("sent message: '%s'\n", message)
			fmt.Println(event)
			counter += 1
		case <-time.After(1 * time.Second):
			log.Fatalf("failed to send message '%s'", message)
		}
	}
}
