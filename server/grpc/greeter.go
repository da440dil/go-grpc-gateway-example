package grpc

import (
	"context"

	"github.com/da440dil/go-grpc-gateway-example/greeter"
	pb "github.com/da440dil/go-grpc-gateway-example/proto"
)

// GreeterServer implements pb.GreeterServer.
type GreeterServer struct {
	greeter *greeter.Greeter
}

// NewGreeterServer creates new GreeterServer.
func NewGreeterServer() *GreeterServer {
	return &GreeterServer{greeter: greeter.NewGreeter()}
}

// SayHello says hello.
func (s *GreeterServer) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloReply, error) {
	v, err := s.greeter.SayHello(r.Name)
	if err != nil {
		return nil, err
	}
	return &pb.HelloReply{Message: v}, nil
}
