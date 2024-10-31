package main

import (
	"flag"
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var KAFKA_BROKER_ADDRESS = os.Getenv("KAFKA_BROKER_ADDRESS")

func main() {
	mode := flag.String("mode", "producer", "sets the mode of example program to run, either 'producer' or 'consumer'.")
	targetTopic := flag.String("targetTopic", "testTopic", "sets the topic to produce/consume messages to")
	flag.Parse()

	// Create Kafka connection configmap
	config := &kafka.ConfigMap{
		// Producer-required configs
		"bootstrap.servers": KAFKA_BROKER_ADDRESS,
		"client.id":         "exampleKafkaClient",
		"acks":              "all",
		// Consumer-required configs
		"auto.offset.reset": "smallest",
		"group.id":          "exampleConsumer",
	}

	log.Printf("starting Kafka example in mode: %s", *mode)
	if *mode == "producer" {
		kafkaProducer(*targetTopic, config)
	} else if *mode == "consumer" {
		kafkaConsumer(*targetTopic, config)
	} else {
		log.Fatalf("unknown mode provided: %s", mode)
	}
}
