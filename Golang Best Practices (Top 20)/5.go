package main

import "fmt"

// namedReturn demonstrates named return values.
func namedReturn(x, y int) (result int) {
	result = x + y
	return
}

// explicitReturn demonstrates explicit return values.
func explicitReturn(x, y int) int {
	return x + y
}

func main() {
	// Named return values
	sum1 := namedReturn(3, 5)
	fmt.Println("Named Return:", sum1)

	// Explicit return values
	sum2 := explicitReturn(3, 5)
	fmt.Println("Explicit Return:", sum2)
}
