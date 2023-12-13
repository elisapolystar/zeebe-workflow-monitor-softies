package main

import (
  "context"
  "fmt"
  "reflect"
  "os"
  "os/signal"
  "time"

  kafka "github.com/segmentio/kafka-go"
)

// Helper function for printing the reader configuration for debugging
func printConfig(config kafka.ReaderConfig) {
  configType := reflect.TypeOf(config)
  configValue := reflect.ValueOf(config)

  for i := 0; i < configType.NumField(); i++ {
    field := configType.Field(i)
    value := configValue.Field(i).Interface()
    fmt.Printf("%s: %v\n", field.Name, value)
  }
}

func Consume(messageChannel chan<- topicMessagePair) {
  // Define the Kafka broker address and topics we want to subscribe to
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

  // Set up a signal channel to handle termination
  signals := make(chan os.Signal, 1)
  signal.Notify(signals, os.Interrupt)

  // Subscribe to each topic concurrently
  for _, topic := range topics {
    go func(t string) {
      // Create a Kafka reader
      r := kafka.NewReader(kafka.ReaderConfig{
        Brokers:   brokers,
        Topic:     t,
        MinBytes:  10e3, // 10KB
        MaxBytes:  10e6, // 10MB
        MaxWait:   250 * time.Millisecond,
        Partition: 0,
        StartOffset: kafka.LastOffset,
        IsolationLevel: 1,
      })

      defer r.Close()

      // Consume messages from the Kafka topics
      for {
        m, err := r.FetchMessage(context.Background())
        if err != nil {
          fmt.Printf("Error fetching message from topic %s: %v\n", t, err)
          break
        }

        //config := r.Config()
        //fmt.Printf("Received tunturi from topic %s: %s\n", m.Topic, string(m.Value))
        //fmt.Printf("message at %s offset %d partition %d: %s = %s\n", m.Topic, m.Offset, m.Partition, string(m.Key), string(m.Value))
        //fmt.Printf("Received message: %+v\n", m)
        //printConfig(config)
        //fmt.Printf("Config: %v\n", config)
        fmt.Printf("len key: %d && len value: %d from topic %s with offset %d\n", len(m.Key), len(m.Value), m.Topic, m.Offset)
        //messageChannel <- topicMessagePair{m.Topic, m.Value}
      }
    }(topic)
  }

  // Block until a signal is received
  <-signals
}
