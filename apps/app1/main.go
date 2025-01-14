package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	// Read environment variables with defaults
	app2Host := os.Getenv("APP2_HOST")
	if app2Host == "" {
		app2Host = "app2" // Default to app2 service name
	}

	app2Port := os.Getenv("APP2_PORT")
	if app2Port == "" {
		app2Port = "8081" // Default port for app2
	}

	// HTTP handler for the root endpoint
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Simple response to show that app1 is working
		fmt.Fprintf(w, "Welcome to app1! It can also make requests to app2.")
	})

	// HTTP handler for making requests to app2
	http.HandleFunc("/request-to-app2", func(w http.ResponseWriter, r *http.Request) {
		// Construct the URL for app2's /hello endpoint
		service2URL := fmt.Sprintf("http://%s:%s/hello", app2Host, app2Port)

		// Make the request to app2
		resp, err := http.Get(service2URL)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error making request to app2: %v", err), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Read the response from app2
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error reading response: %v", err), http.StatusInternalServerError)
			return
		}

		// Write the response from app2 to the client
		fmt.Fprintf(w, "Response from app2: %s", body)
	})

	// Start the server on port 8080 for app1
	port := os.Getenv("APP1_PORT")
	if port == "" {
		port = "8080" // Default port for app1
	}

	fmt.Printf("app1 listening on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Error starting app1: %v", err)
	}
}