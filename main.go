package main

import (
	"fmt"
	"sync"

	"github.com/chnmk/grpc-rest-concurrency/counter"
	grpc_api "github.com/chnmk/grpc-rest-concurrency/grpc"
	rest_api "github.com/chnmk/grpc-rest-concurrency/rest"
)

func main() {
	go rest_api.Server()
	go grpc_api.Server()

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
			grpc_api.Client()
		}
	}()

	for i := 0; i < 50; i++ {
		ch <- i
	}

	close(ch)

	wg.Wait()

	fmt.Println("REST: ", counter.Counter_REST)
	fmt.Println("gRPC: ", counter.Counter_gRPC)
}
