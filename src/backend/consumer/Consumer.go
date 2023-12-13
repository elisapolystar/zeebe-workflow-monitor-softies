package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/IBM/sarama"
)

// Subscribe to wanted kafka topics, listen to messages from these topics and pass the messages with their corresponding topics to a channel
func Consume(messageChannel chan<- topicMessagePair) {

	// Define the Kafka broker address and topic we want to subscribe to
	brokers := []string{broker}
	topics := []string{zeebe,
		zeebe_deployment,
		zeebe_deploy_distribution,
		zeebe_error,
		zeebe_incident,
		zeebe_job_batch,
		zeebe_job,
		zeebe_message,
		zeebe_message_subscription,
		zeebe_message_subscription_start_event,
		zeebe_process,
		zeebe_process_event,
		zeebe_process_instance,
		zeebe_process_instance_result,
		zeebe_process_message_subscription,
		zeebe_timer,
		zeebe_variable}

	// Configure the Kafka consumer
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// Create a Kafka consumer that will be used to create partition consumers for each topic
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

	// Create a map to store partition consumers for each topic, topic is the key and the partition consumer is the value
	partitionConsumers := make(map[string]sarama.PartitionConsumer)

	// Subscribe to each topic declared in the topics list
	for _, topic := range topics {

		// Create a partition consumer for a specific topic
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

		// Add the partition consumer of a specific topic to the map with the corresponding topic as the key
		partitionConsumers[topic] = partitionConsumer
		fmt.Printf("Subscribed to topic: %s\n", topic)
	}

	// Infinite loop that listens to all the topics in the partitionConsumers map
	// If we get a message, the message and the topic are passed to the channel
	for {
		select {
		// If there is a termination signal
		case <-signals:
			fmt.Println("Received termination signal. Closing consumer.")
			return
		default:
			// Loop the consumers of the topics if they have messages
			for topic, partitionConsumer := range partitionConsumers {
				select {
				// Receive a message from some of the partition consumers
				case msg := <-partitionConsumer.Messages():

					// Put the received message and the topic to the channel
					messageChannel <- topicMessagePair{topic, msg.Value}
				case err := <-partitionConsumer.Errors():
					fmt.Printf("Error consuming message from topic %s: %v\n", topic, err)

				// If there are no messages or errors from this topic we continue to the next item in the for-loop
				default:
					continue
				}
			}
		}
		// Pause for 0.1 seconds to save cpu usage
		time.Sleep(100 * time.Millisecond)
	}
}
