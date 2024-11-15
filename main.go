package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/chnmk/grpc-rest-concurrency/counter"
	grpc_api "github.com/chnmk/grpc-rest-concurrency/grpc"
	pb "github.com/chnmk/grpc-rest-concurrency/grpc/example"
	rest_api "github.com/chnmk/grpc-rest-concurrency/rest"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Start the servers in goroutines
	go rest_api.Server()
	go grpc_api.Server()

	// Create a gRPC client
	conn, err := grpc.NewClient("127.0.0.1:50051", // May not work properly with "localhost" address
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Wait for the servers to start
	time.Sleep(time.Second * 5)

	// Create clients to read signals from a channel
	var wg sync.WaitGroup

	ch := make(chan int, 5)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for range ch {
			rest_api.Client()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for range ch {
			grpc_api.Client(c)
		}
	}()

	// Send signals to a channel
	for i := 0; i < 50; i++ {
		ch <- i
	}

	close(ch)

	wg.Wait()

	// Check which technology handled more signals
	fmt.Println("REST: ", counter.Counter_REST)
	fmt.Println("gRPC: ", counter.Counter_gRPC)
}
