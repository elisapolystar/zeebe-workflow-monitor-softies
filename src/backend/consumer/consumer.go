package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/IBM/sarama"
	"github.com/gorilla/websocket"
)

func main() {

	// Define the Kafka broker address and topic we want to subscribe to
	brokers := []string{"kafka:9092"}
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

	// Create a channel to send messages to the frontend
	messageChannel := make(chan string)

	// Consume messages from the Kafka topics
	for {
		select {
		case topic := <-partitionConsumers:
			select {
			case msg := <-partitionConsumers[topic].Messages():
				message := string(msg.Value)
				fmt.Printf("Received message from topic %s: %s\n", topic, string(msg.Value))
				messageChannel <- message
			case err := <-partitionConsumers[topic].Errors():
				fmt.Printf("Error consuming message from topic %s: %v\n", topic, err)
			case <-signals:
				fmt.Println("Received termination signal. Closing consumer.")
				return
			}
		}
	}

	// Handle HTTP requests from the frontend
	http.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {
		// Set the response content type to JSON
		w.Header().Set("Content-Type", "application/json")

		// Read messages from channel and send them to frontend
		message := <-messageChannel

		// Create a JSON response
		response := map[string]string{"message": message}

		// Serialize the reponse to Json
		jsonBytes, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Write JSON to the frontend
		w.WriteHeader(http.StatusOK)
		w.Write(jsonBytes)
	})

	// Establish a WebSocket connection to the frontend
	// (Server URL will be updated once it's known)
	serverURL := "URL"
	test_message := "Hello from the backend"
	conn, _, err := websocket.DefaultDialer.Dial(serverURL, nil)
	if err != nil {
		log.Fatal("Error connecting to WebSocket:", err)
	}
	defer conn.Close()

	// Send test message to the frontend
	err := conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		log.Println("Error sending message:", err)
		return
	}
	else {
		fmt.Println("Sent message:", message)
	}	
}