package main

import "fmt"

// Option struct to hold configuration options
type Option struct {
	Port    int
	Timeout int
}

// ServerConfig is a function that accepts an Option struct
func ServerConfig(opt Option) {
	fmt.Printf("Server configuration - Port: %d, Timeout: %d seconds\n", opt.Port, opt.Timeout)
}

func main() {
	// Creating an Option struct with default values
	defaultConfig := Option{
		Port:    8080,
		Timeout: 30,
	}

	// Configuring the server with default options
	ServerConfig(defaultConfig)

	// Modifying the Port using a new Option struct
	customConfig := Option{
		Port: 9090,
	}

	// Configuring the server with custom Port value and default Timeout
	ServerConfig(customConfig)
}
