Named return values are commonly used in Go, but they can sometimes make code less clear, especially in larger codebases.


Let’s see the difference using a simple example.


```go
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
```

In the sample program above, we have two functions, namedReturn and explicitReturn. Here's how they differ:


- namedReturn uses a named return value result. Although it's clear what the function is returning, it may not be immediately obvious in more complex functions.
  
- explicitReturn returns the result directly. This is simpler and more explicit.