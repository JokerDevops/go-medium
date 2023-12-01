package main

import "fmt"

func main() {
	// Declare and initialize an outer variable 'x' with the value 10.
	x := 10
	fmt.Println("Outer x:", x)

	// Enter an inner scope with a new variable 'x' shadowing the outer 'x'.
	if true {
		x := 5                     // Shadowing occurs here
		fmt.Println("Inner x:", x) // Print the inner 'x', which is 5.
	}

	// The outer 'x' remains unchanged and is still accessible.
	fmt.Println("Outer x after inner scope:", x) // Print the outer 'x', which is 10.
}
