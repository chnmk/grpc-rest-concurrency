package grpc_api

import (
	"context"
	"log"
	"time"

	pb "github.com/chnmk/grpc-rest-concurrency/grpc/example"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Client() {
	// May not work properly with "localhost" address
	conn, err := grpc.NewClient("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = c.SayHello(ctx, &pb.HelloRequest{Name: "World"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	// log.Printf("Greeting: %s", r.GetMessage())
}
