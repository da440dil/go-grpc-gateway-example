package server

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	srvg "github.com/da440dil/go-grpc-gateway-example/server/grpc"
	srvh "github.com/da440dil/go-grpc-gateway-example/server/http"
)

// ServeGRPC returns function which starts gRPC server.
func ServeGRPC(addr string, stopTimeout time.Duration) func(<-chan struct{}) error {
	return func(stop <-chan struct{}) error {
		srv := srvg.NewServer()
		done := make(chan error, 2)
		go func() {
			<-stop
			log.Println("GRPC server is about to stop")
			done <- srv.Stop(stopTimeout)
		}()

		go func() {
			log.Printf("GRPC server starts listening at %s\n", addr)
			done <- srv.Start(addr)
		}()

		for i := 0; i < cap(done); i++ {
			if err := <-done; err != nil {
				log.Printf("GRPC server stopped with error %v\n", err)
				return err
			}
		}
		log.Println("GRPC server stopped")
		return nil
	}
}

// ServeHTTP returns function which starts HTTP server.
func ServeHTTP(addr, grpcAddr string, stopTimeout time.Duration) func(<-chan struct{}) error {
	return func(stop <-chan struct{}) error {
		srv := srvh.NewServer()
		done := make(chan error, 2)
		go func() {
			<-stop
			log.Println("HTTP server is about to stop")
			done <- srv.Stop(stopTimeout)
		}()

		go func() {
			log.Printf("HTTP server starts listening at %s\n", addr)
			done <- srv.Start(addr, grpcAddr)
		}()

		for i := 0; i < cap(done); i++ {
			if err := <-done; err != nil && err != http.ErrServerClosed {
				log.Printf("HTTP server stopped with error %v\n", err)
				return err
			}
		}
		log.Println("HTTP server stopped")
		return nil
	}
}

// ListenOsSignal returns function which starts listening os signal.
func ListenOsSignal(v ...os.Signal) func(<-chan struct{}) error {
	return func(stop <-chan struct{}) error {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, v...)
		go func() {
			<-stop
			close(sig)
		}()

		log.Println("Server starts listening os signal")
		<-sig
		log.Println("Server stops listening os signal")
		signal.Stop(sig)
		return nil
	}
}
