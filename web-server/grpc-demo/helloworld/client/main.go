package main

import (
	"context"
	"log"
	"time"

	"github.com/wgzhulin/golang-study-with-test/web-server/grpc-demo/helloworld/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("not connect: %v", err)
	}

	defer conn.Close()

	c := helloworld.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &helloworld.HelloRequest{Name: "wangzhulin"})
	if err != nil {
		log.Fatal("call fail: %v", err)
	}

	log.Printf("Greeting: %s", r.GetMessage())

	rAgain, err := c.SayHelloAgain(ctx, &helloworld.HelloRequest{Name: "wangzhulin"})
	if err != nil {
		log.Fatal("call fail: %v", err)
	}

	log.Printf("Greeting: %s", rAgain.GetMessage())
}
