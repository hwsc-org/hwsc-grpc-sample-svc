package service

import (
	"fmt"
	pb "github.com/faraonc/hwsc-api-blocks/int/hwsc-grpc-sample-svc/proto"
	"github.com/golang/protobuf/ptypes"
	"golang.org/x/net/context"
	"time"
)

type Service struct{}

// This matches the proto in https://github.com/faraonc/hwsc-api-blocks/blob/master/int/hwsc-grpc-sample-svc/proto/grpc-sample-svc.proto
func (s Service) SayHello(ctx context.Context, in *pb.SampleServiceRequest) (*pb.SampleServiceResponse, error) {
	fmt.Printf("Request { %v}\n", in)
	fmt.Printf("Request date: %s\n", time.Unix(in.CreateTimestamp.GetSeconds(),
		int64(in.CreateTimestamp.GetNanos())))

	// Use Google Protobuf Timestamp
	t := ptypes.TimestampNow()

	response := &pb.SampleServiceResponse{
		Message: "Hello " + in.FirstName,
		// TODO Convert to our own timestamp until resolve
		// https://github.com/faraonc/hwsc-gateway-svc/issues/37
		ResponseTimestamp: &pb.Timestamp{
			Seconds: t.GetSeconds(),
			Nanos:   t.GetNanos(),
		},
	}
	return response, nil
}
