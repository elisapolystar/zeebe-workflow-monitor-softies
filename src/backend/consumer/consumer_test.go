package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestConsumer(t *testing.T) {
	t.Run("Kafka not running", func(t *testing.T) {
		// Redirect os.Stdout to a pipe
		oldStdout := os.Stdout
		defer func() { os.Stdout = oldStdout }()
		r, w, _ := os.Pipe()
		os.Stdout = w

		// Call the function that prints the error
		messageChannel = make(chan string)
		Consume(messageChannel)

		// Close the write end of the pipe to release the resources
		w.Close()

		// Read from the read end of the pipe
		var buf bytes.Buffer
		buf.ReadFrom(r)

		// Check if the printed error message starts with the expected prefix
		expectedErrorMessage := "Error creating Kafka consumer:"
		if !strings.HasPrefix(buf.String(), expectedErrorMessage) {
			t.Errorf("Expected error message to start with '%s', but got '%s'", expectedErrorMessage, buf.String())
		}
	})
}
