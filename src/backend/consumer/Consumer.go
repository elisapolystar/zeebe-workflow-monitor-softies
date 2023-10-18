package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/IBM/sarama"
)

func Consume(messageChannel chan<- string) {

	// Define the Kafka broker address and topic we want to subscribe to
	brokers := []string{"kafka:9093"}
	topics := []string{"zeebe",
		"zeebe-deployment",
		"zeebe-deploy-distribution",
		"zeebe-error",
		"zeebe-incident",
		"zeebe-job-batch",
		"zeebe-job",
		"zeebe-message",
		"zeebe-message-subscription",
		"zeebe-message-subscription-start-event",
		"zeebe-process",
		"zeebe-process-event",
		"zeebe-process-instance",
		"zeebe-process-instance-result",
		"zeebe-process-message-subscription",
		"zeebe-timer",
		"zeebe-variable"}

	// Configure the Kafka consumer
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// Create a Kafka consumer
	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		fmt.Printf("Error creating Kafka consumer: %v\n", err)
		return
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			fmt.Printf("Error closing Kafka consumer: %v\n", err)
		}
	}()

	// Set up a signal channel to handle termination
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// Create a map to store partition consumers for each topic
	partitionConsumers := make(map[string]sarama.PartitionConsumer)

	// Subscribe to each topic
	for _, topic := range topics {
		partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
		if err != nil {
			fmt.Printf("Error subscribing to topic %s: %v\n", topic, err)
			return
		}
		defer func(topic string) {
			if err := partitionConsumers[topic].Close(); err != nil {
				fmt.Printf("Error closing partition consumer for topic %s: %v\n", topic, err)
			}
		}(topic)

		partitionConsumers[topic] = partitionConsumer
		fmt.Printf("Subscribed to topic: %s\n", topic)
	}

	// Consume messages from the Kafka topics
	for {
		for topic, partitionConsumer := range partitionConsumers {
			select {
			case msg := <-partitionConsumer.Messages():
				fmt.Printf("Received message from topic %s: %s\n", topic, string(msg.Value))
				messageChannel <- string(msg.Value)
			case err := <-partitionConsumer.Errors():
				fmt.Printf("Error consuming message from topic %s: %v\n", topic, err)
			default:
				continue
			}
		}
	}
}
