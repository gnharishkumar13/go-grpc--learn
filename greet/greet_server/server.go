package main

import (
	"fmt"
	"github.com/gnharishkumar13/go-grpc--learn/greet/greetpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {}

func main() {
	fmt.Println("Hello from server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(s, &server{})
	fmt.Println("server started")
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve %v", err)
	}


}
