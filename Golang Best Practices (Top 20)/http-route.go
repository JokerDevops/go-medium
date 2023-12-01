package main

import (
	"fmt"
	"net/http"
)

// Router defines the interface for a router.
type Router interface {
	HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
	ServeHTTP(http.ResponseWriter, *http.Request)
}

// MyRouter is an example implementation of the Router interface.
type MyRouter struct {
	routes map[string]func(http.ResponseWriter, *http.Request)
}

func (r *MyRouter) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	r.routes[pattern] = handler
}

func (r *MyRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	handler, ok := r.routes[req.URL.Path]
	if !ok {
		http.NotFound(w, req)
		return
	}
	handler(w, req)
}

// LoggerMiddleware is an example middleware.
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Logging:", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
func NewMyRouter() *MyRouter {
	return &MyRouter{routes: make(map[string]func(http.ResponseWriter, *http.Request))}
}
func main() {
	// Create a router instance
	router := NewMyRouter()
	// Attach middleware
	http.Handle("/", LoggerMiddleware(router))

	// Define routes
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to our website!")
	})

	router.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "About Us")
	})

	router.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Contact Us")
	})

	// Start the server
	fmt.Println("Server is listening on :8080")
	http.ListenAndServe(":8080", nil)
}
