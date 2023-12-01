package main

import (
	"fmt"
	"time"
)

func main() {
	// create a string slice
	data := []string{"one", "two", "three"}

	// loop through the slice
	for i, v := range data {
		// create a goroutine to print the values
		go func() {
			// print the values
			fmt.Println(i, v)
		}()
	}

	// wait 3 seconds
	time.Sleep(3 * time.Second)
	// goroutines print: three, three, three
}
