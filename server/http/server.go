package http

import (
	"context"
	"net/http"
	"time"

	pb "github.com/da440dil/go-grpc-gateway-example/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

// Server contains http.Server.
type Server struct {
	server *http.Server
}

// NewServer creates new Server.
func NewServer() *Server {
	return &Server{server: &http.Server{}}
}

// Start starts http.Server.
func (s *Server) Start(addr, grpcAddr string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterGreeterHandlerFromEndpoint(ctx, gwmux, grpcAddr, opts)
	if err != nil {
		return err
	}

	s.server.Addr = addr
	s.server.Handler = gwmux
	return s.server.ListenAndServe()
}

// Stop stops http.Server.
func (s *Server) Stop(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
