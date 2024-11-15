package grpc_api

import (
	"context"
	"log"
	"time"

	pb "github.com/chnmk/grpc-rest-concurrency/grpc/example"
)

func Client(c pb.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := c.SayHello(ctx, &pb.HelloRequest{Name: "World"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	// log.Printf("Greeting: %s", r.GetMessage())
}
