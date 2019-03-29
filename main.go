package main

import (
	"log"
	"os"
	"time"

	"github.com/da440dil/go-grpc-gateway-example/server"
	"github.com/da440dil/go-workgroup"
)

func main() {
	grpcAddr := "127.0.0.1:50051"
	stopTimeout := time.Second * 10

	var wg workgroup.Group
	wg.Add(server.ServeGRPC(grpcAddr, stopTimeout))
	wg.Add(server.ListenOsSignal(os.Interrupt))
	if err := wg.Run(); err != nil {
		log.Fatal(err)
	}
}
