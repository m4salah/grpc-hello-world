package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/m4salah/grpc-hello-world/helloworld"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type helloWorldServer struct {
	pb.UnimplementedHelloWorldServer
}

func (h helloWorldServer) HelloWorld(context.Context, *pb.HelloWorldRequest) (*pb.HelloWorldResponse, error) {
	return &pb.HelloWorldResponse{Message: "Hello, World!"}, nil
}
func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloWorldServer(s, &helloWorldServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
