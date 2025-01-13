

package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"

	"google.golang.org/grpc"
	pb "./proto"
)

type server struct {
	pb.UnimplementedBackendServiceServer
}

func (s *server) ProcessRequest(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	log.Printf("Received gRPC request with message: %s", req.Message)
	return &pb.Response{Response: "Processed gRPC request successfully!"}, nil
}

func main() {
	http.HandleFunc("/api/process", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received request at /api/process")
		response := "Processed REST request successfully!"
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	grpcPort := ":50051"
	listener, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", grpcPort, err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterBackendServiceServer(grpcServer, &server{})

	go func() {
		log.Printf("App B is running gRPC server on port %s", grpcPort)
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()

	log.Printf("App B is running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}