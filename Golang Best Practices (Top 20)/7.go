package main

import (
	"fmt"
)

// Define a struct type representing a person
type Person struct {
	FirstName string // First name of the person
	LastName  string // Last name of the person
	Age       int    // Age of the person
}

func main() {
	// Using a composite literal to create a Person instance
	person := Person{
		FirstName: "John", // Initialize the FirstName field
		LastName:  "Doe",  // Initialize the LastName field
		Age:       30,     // Initialize the Age field
	}

	// Printing the person's information
	fmt.Println("Person Details:")
	fmt.Println("First Name:", person.FirstName) // Access and print the First Name field
	fmt.Println("Last Name:", person.LastName)   // Access and print the Last Name field
	fmt.Println("Age:", person.Age)              // Access and print the Age field
}
