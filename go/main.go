package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var KAFKA_BROKER_ADDRESS = os.Getenv("KAFKA_BROKER_ADDRESS")

func main() {
	mode := flag.String("mode", "producer", "sets the mode of example program to run, either 'producer' or 'consumer'.")
	targetTopic := flag.String("targetTopic", "testTopic", "sets the topic to produce/consume messages to")
	waitTime := flag.Int("wait", 0, "waits desired number of seconds before trying to reach Kafka")
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

	fmt.Printf("starting Kafka example in mode: %s\n", *mode)
	if *waitTime != 0 {
		fmt.Printf("waiting for %d seconds before trying to reach Kafka\n", *waitTime)
		<-time.After(time.Duration(*waitTime) * time.Second)
	}

	if *mode == "producer" {
		kafkaProducer(*targetTopic, config)
	} else if *mode == "consumer" {
		kafkaConsumer(*targetTopic, config)
	} else {
		log.Fatalf("unknown mode provided: %s", mode)
	}
}
