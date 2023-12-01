// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	// Open a file
// 	file, err := os.Open("/Users/jameszhang/Project/go-medium/Golang Best Practices (Top 20)/example.txt")
// 	if err != nil {
// 		// Handle the error
// 		fmt.Println("Error opening the file:", err)
// 		return
// 	}
// 	defer file.Close() // Close the file when done

// 	// Read from the file
// 	buffer := make([]byte, 1024)
// 	_, err = file.Read(buffer)
// 	if err != nil {
// 		// Handle the error
// 		fmt.Println("Error reading the file:", err)
// 		return
// 	}

// 	// Print the file content
// 	fmt.Println("File content:", string(buffer))
// }

// #12: Comment Your Code

// Function Comments
// Add comments to functions to explain their purpose, parameters, and return values. Use the `godoc` style for function comments.
// package main

// greetUser greets a user by name.
// Parameters:
//
//	name (string): The name of the user to greet.
//
// Returns:
//
//	string: The greeting message.
// func greetUser(name string) string {
// 	return "Hello, " + name + "!"
// }

// func main() {
// 	userName := "Alice"
// 	greeting := greetUser(userName)
// 	fmt.Println(greeting)
// }

// package main

// import "fmt"

// // This is the main package of our Go program.
// // It contains the entry point (main) function.
// func main() {
//     fmt.Println("Hello, World!")
// }