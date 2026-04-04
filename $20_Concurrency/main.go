package main

import "fmt"

func main() {
	// Create a string type channel
	messages := make(chan string)

	// Run an anonymous concurrent function
	go func() {
		// Send data to channel "messages"
		messages <- "Hello from Goroutine!"
	}()

	// Import data from channel "message
	msg := <-messages
	fmt.Println(msg)
}
