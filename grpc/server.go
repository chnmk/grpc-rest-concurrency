package grpc_api

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"sync/atomic"

	"github.com/chnmk/grpc-rest-concurrency/counter"
	pb "github.com/chnmk/grpc-rest-concurrency/grpc/example"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(_ context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	// log.Printf("Received: %v", in.GetName())
	if in.GetName() == "World" {
		atomic.AddInt32(&counter.Counter_gRPC, 1)
	}
	return &pb.HelloReply{Message: "Hello " + in.GetName() + "!"}, nil
}

func Server() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
