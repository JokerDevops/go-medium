package main

import "fmt"

// Function that might panic
func riskyOperation() {
	defer func() {
		if r := recover(); r != nil {
			// Recover from the panic and handle it gracefully
			fmt.Println("Recovered from panic:", r)
		}
	}()

	// Simulate a panic condition
	panic("Oops! Something went wrong.")
}

func main() {
	fmt.Println("Start of the program.")

	// Call the risky operation within a function that recovers from panics
	riskyOperation()

	fmt.Println("End of the program.")
}
