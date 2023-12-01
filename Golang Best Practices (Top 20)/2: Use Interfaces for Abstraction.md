Abstraction
Abstraction is a fundamental concept in Go, allowing us to define behavior without specifying implementation details.

Interfaces
In Go, an interface is a collection of method signatures.

Any type that implements all the methods of an interface implicitly satisfies that interface.

This enables us to write code that can work with different types as long as they adhere to the same interface.

Here’s a single sample program in Go that demonstrates the concept of using interfaces for abstraction:

```go
package main

import (
    "fmt"
    "math"
)

// Define the Shape interface
type Shape interface {
    Area() float64
}

// Rectangle struct
type Rectangle struct {
    Width  float64
    Height float64
}

// Circle struct
type Circle struct {
    Radius float64
}

// Implement the Area method for Rectangle
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Implement the Area method for Circle
func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

// Function to print the area of any Shape
func PrintArea(s Shape) {
    fmt.Printf("Area: %.2f\n", s.Area())
}

func main() {
    rectangle := Rectangle{Width: 5, Height: 3}
    circle := Circle{Radius: 2.5}

    // Call PrintArea on rectangle and circle, both of which implement the Shape interface
    PrintArea(rectangle) // Prints the area of the rectangle
    PrintArea(circle)    // Prints the area of the circle
}
```

In this single program, we define the Shape interface, create two structs Rectangle and Circle, each implementing the Area() method, and use the PrintArea function to print the area of any shape that satisfies the Shape interface.

This demonstrates how you can use interfaces for abstraction in Go to work with different types using a common interface.

