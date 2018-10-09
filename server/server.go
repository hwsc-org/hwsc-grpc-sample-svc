package server

import(
	"golang.org/x/net/context"

	pb"github.com/faraonc/hwsc-api-blocks/int/hwsc-grpc-sample/proto"
)

type Server struct{}

// This matches the proto in https://github.com/faraonc/hwsc-api-blocks/blob/master/int/hwsc-grpc-sample/proto/grpc-sample.proto
func (s Server) SayHello(ctx context.Context, in *pb.SampleServiceRequest) (*pb.SampleServiceResponse, error) {
	return &pb.SampleServiceResponse{Message: "Hello " + in.Name}, nil
}