package grpc

import (
	"context"
	"net"
	"time"

	pb "github.com/da440dil/go-grpc-gateway-example/proto"
	"google.golang.org/grpc"
)

// Server contains grpc.Server.
type Server struct {
	server *grpc.Server
}

// NewServer creates new Server.
func NewServer() *Server {
	return &Server{server: grpc.NewServer()}
}

// Start starts grpc.Server.
func (s *Server) Start(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	srv := NewGreeterServer()
	pb.RegisterGreeterServer(s.server, srv)
	return s.server.Serve(lis)
}

// Stop stops grpc.Server.
func (s *Server) Stop(timeout time.Duration) error {
	done := make(chan struct{})
	go func() {
		s.server.GracefulStop()
		close(done)
	}()

	select {
	case <-done:
		return nil
	case <-time.After(timeout):
		s.server.Stop()
		return context.DeadlineExceeded
	}
}
