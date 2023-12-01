package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(10 * time.Second)
		w.Write([]byte("Hello, World!"))
	})

	server := &http.Server{Addr: ":18080"}

	stopChan := make(chan os.Signal, 1)

	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM, os.Kill)

	// Start the HTTP server in a goroutine
	go func() {
		// If ListenAndServe returns an error and it's not a server closed error,
		// then log it as a fatal error.
		log.Println("Server is startting...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %s", err)
		}
	}()

	// Wait until we get a stop signal
	sig := <-stopChan

	// Log that the server is shutting down
	log.Println("Server is shutting down, due to:", sig)

	// Create a context with a 15-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Initiate graceful shutdown
	// if it doesn't complete in 15 seconds, it will be forcefully stopped
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %s", err)
	} else {
		log.Println("Server stopped gracefully")
	}
	close(stopChan)

}
