// App A: Frontend Service
package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"google.golang.org/grpc"
	pb "./proto"
)

//const backendServiceURL = "http://app-b:8080/api/process"
//const grpcServiceAddress = "app-b:50051"

var (
    backendServiceURL string = os.Getenv("backendServiceURL")
    grpcServiceAddress string = os.Getenv("grpcServiceAddress")
)


func main() {
	http.HandleFunc("/api/trigger", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received request at /api/trigger")

		resp, err := http.Get(backendServiceURL)
		if err != nil {
			log.Printf("Error calling backend service: %v", err)
			http.Error(w, "Failed to call backend service", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		w.WriteHeader(resp.StatusCode)
		w.Write([]byte("REST call to backend succeeded."))
	})

	http.HandleFunc("/api/grpc-trigger", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received request at /api/grpc-trigger")
		conn, err := grpc.Dial(grpcServiceAddress, grpc.WithInsecure())
		if err != nil {
			log.Printf("Failed to connect to gRPC server: %v", err)
			http.Error(w, "Failed to connect to gRPC service", http.StatusInternalServerError)
			return
		}
		defer conn.Close()

		client := pb.NewBackendServiceClient(conn)
		response, err := client.ProcessRequest(context.Background(), &pb.Request{Message: "Trigger from App A"})
		if err != nil {
			log.Printf("Error calling gRPC service: %v", err)
			http.Error(w, "Failed to call gRPC service", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response.Response))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("App A is running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}