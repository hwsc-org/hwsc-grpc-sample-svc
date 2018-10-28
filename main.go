package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/faraonc/hwsc-api-blocks/int/hwsc-grpc-sample-svc/proto"
	svc "github.com/faraonc/hwsc-grpc-sample-svc/service"
)

func main() {
	log.Println("[INFO] hwsc-grpc-sample-svc initiating...")

	// Step 1: Make TCP listener
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("[FATAL] Failed to initialize TCP listener %v\n", err)
	}

	// Step 2: Make gRPC server
	s := grpc.NewServer()

	// Step 3: Implement a service in a folder service/service.go
	// Step 4: Register the service with gRPC server
	pb.RegisterSampleServiceServer(s, svc.Service{})
	log.Println("[INFO] hwsc-grpc-sample-svc at localhost:50051...")

	// Step 5: Start gRPC server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
