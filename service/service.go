package service

import (
	"golang.org/x/net/context"

	pb "github.com/faraonc/hwsc-api-blocks/int/hwsc-grpc-sample-svc/proto"
)

type Service struct{}

// This matches the proto in https://github.com/faraonc/hwsc-api-blocks/blob/master/int/hwsc-grpc-sample-svc/proto/grpc-sample-svc.proto
func (s Service) SayHello(ctx context.Context, in *pb.SampleServiceRequest) (*pb.SampleServiceResponse, error) {
	return &pb.SampleServiceResponse{Message: "Hello " + in.FirstName}, nil
}
