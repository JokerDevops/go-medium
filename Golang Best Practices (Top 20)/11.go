package main

import (
	"fmt"
	"time"
)

// Function that runs concurrently
func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(100 * time.Millisecond)
	}
}

// Function that runs in the main goroutine
func main() {
	// Start the goroutine
	go printNumbers()

	// Continue executing main
	for i := 0; i < 2; i++ {
		fmt.Println("Hello")
		time.Sleep(200 * time.Millisecond)
	}
	// Ensure the goroutine completes before exiting
	time.Sleep(1 * time.Second)
}
