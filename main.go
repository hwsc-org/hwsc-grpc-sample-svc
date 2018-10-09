package main

import (
	"fmt"
	"github.com/faraonc/hwsc-api-blocks/int/hwsc-grpc-sample/proto"
)

func main() {
	fmt.Println("Hello, playground")
	grpc_sample_pb.RegisterSampleServiceServer()
}
