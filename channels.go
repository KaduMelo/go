package main

import "fmt"

// Routine 1
func main() {
	canal := make(chan string)

	// Routine 2
	go func() {
		canal <- "Arriving messages - Coming from routine 2!"
	}()

	// Routine 1
	message := <-canal // if channel is full, empty
	fmt.Println(message)
}
