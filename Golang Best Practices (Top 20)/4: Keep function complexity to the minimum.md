Function complexity refers to the degree of intricacy, nesting, and branching within a function’s code. Keeping function complexity low makes your code more readable, maintainable, and less prone to errors.

Let’s explore this concept with a simple example:

```go
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
```

In above sample program:

We define two functions, CalculateSum and PrintSum, with specific responsibilities.
CalculateSum is a simple function that calculates the sum of two numbers.
PrintSum uses CalculateSum to calculate and print the sum of 5 and 3.
By keeping functions concise and focused on a single task, we maintain low function complexity, improving code readability and maintainability.