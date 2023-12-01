package main

import (
	"fmt"
	"os"
)

func main() {
	// Open the file (Replace "example.txt" with your file's name)
	// file, err := os.Open("/Users/jameszhang/Project/go-medium/Golang Best Practices (Top 20)/example.txt")
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return // Exit the program on error
	}
	defer file.Close() // Ensure the file is closed when the function exits

	// Read and print the contents of the file
	data := make([]byte, 100)
	n, err := file.Read(data)
	if err != nil {
		fmt.Println("Error reading the file:", err)
		return // Exit the program on error
	}

	fmt.Printf("Read %d bytes: %s\n", n, data[:n])
}
