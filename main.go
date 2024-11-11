package main

import (
	"fmt"
	"sync"

	"github.com/chnmk/grpc-rest-concurrency/counter"
	rest_api "github.com/chnmk/grpc-rest-concurrency/rest"
)

func grpc_server() {

}

func grpc_client() {

}

func main() {
	go rest_api.Server()
	go grpc_server()

	var wg sync.WaitGroup

	ch := make(chan int, 5)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for range ch {
			rest_api.Client()
		}
	}()

	/*
		go func() {
			defer wg.Done()
			for range ch {
				grpc_client()
			}
		}()
	*/

	for i := 0; i < 15; i++ {
		ch <- i
	}

	close(ch)

	wg.Wait()

	fmt.Println("REST: ", counter.Counter_REST)
	fmt.Println("gRPC: ", counter.Counter_gRPC)
}
