package main

import (
	"context"
	"log"

	pb "github.com/da440dil/go-grpc-rest-example/proto"
	"google.golang.org/grpc"
)

func main() {
	addr := "127.0.0.1:50051"
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	r, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "World"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("SayHello: { Message: %v }", r.Message)
}
