package main

import (
	"fmt"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Simple response to app1's request
	fmt.Fprintf(w, "Hello from app2!")
}

func main() {
	// Read the port from the environment variable with a default value
	port := os.Getenv("APP2_PORT")
	if port == "" {
		port = "8081" // Default port
	}

	// Handle requests to the /hello endpoint
	http.HandleFunc("/hello", helloHandler)

	// Start the server on the specified port
	fmt.Printf("app2 listening on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Error starting app2 server: %v", err)
	}
}
