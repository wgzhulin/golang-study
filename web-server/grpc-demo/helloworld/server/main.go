package main

import (
	"context"
	"log"
	"net"

	"github.com/wgzhulin/golang-study-with-test/web-server/grpc-demo/helloworld/helloworld"
	"google.golang.org/grpc"
)

type server struct {
	helloworld.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &helloworld.HelloReply{Message: "Hello again " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("faild to listen: %v", err)
	}

	g := grpc.NewServer()

	helloworld.RegisterGreeterServer(g, &server{})

	if err := g.Serve(lis); err != nil {
		log.Fatalf("faild to server: %v", err)
	}
}
