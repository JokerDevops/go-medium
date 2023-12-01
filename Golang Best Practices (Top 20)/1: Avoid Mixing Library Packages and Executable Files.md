In Go, it’s crucial to maintain a clear separation between packages and executable files to ensure clean and maintainable code.

Here’s sample project structure which demonstrates the separation of library and executable file:




```
myproject/
    ├── main.go
    ├── myutils/
       └── myutils.go
```

myutils/myutils.go:

```go 
// Package declaration - Create a separate package for utility functions
package myutils

import "fmt"

// Exported function to print a message
func PrintMessage(message string) {
 fmt.Println("Message from myutils:", message)
}
```

main.go:

```go
// Main program
package main

import (
 "fmt"
 "myproject/myutils" // Import the custom package
)

func main() {
 message := "Hello, Golang!"

 // Call the exported function from the custom package
 myutils.PrintMessage(message)

 // Demonstrate the main program logic
 fmt.Println("Message from main:", message)
}
```


1. In the above example, we have two separate files: myutils.go and main.go.
2. myutils.go defines a custom package named myutils. It contains an exported function PrintMessage that prints a message.
3. main.go is the executable file that imports the custom package myutils using its relative path ("myproject/myutils").
4. The main function in main.go calls the PrintMessage function from the myutils package and prints a message. This separation of concerns keeps the code organized and maintainable.