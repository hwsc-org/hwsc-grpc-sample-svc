package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/faraonc/hwsc-api-blocks/int/hwsc-grpc-sample-svc/proto"
	svr "github.com/faraonc/hwsc-grpc-sample-svc/service"
)

func main() {
	fmt.Println("hwsc-grpc-sample-svc is running...")

	// Step 1: Make TCP listener
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to TCP listener %v", err)
	}

	// Step 2: Make gRPC server
	s := grpc.NewServer()

	// Step 3: Implement a server in a folder /server/server.go
	// Step 4: Register gRPC server
	pb.RegisterSampleServiceServer(s, svr.Service{})

	// Step 5: Start gRPC server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
