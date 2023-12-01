package main

import (
	"fmt"
)

// CalculateSum returns the sum of two numbers.
func CalculateSum(a, b int) int {
	return a + b
}

// PrintSum prints the sum of two numbers.
func PrintSum() {
	x := 5
	y := 3
	sum := CalculateSum(x, y)
	fmt.Printf("Sum of %d and %d is %d\n", x, y, sum)
}

func main() {
	// Call the PrintSum function to demonstrate minimal function complexity.
	PrintSum()
}
