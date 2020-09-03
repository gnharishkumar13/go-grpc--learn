package main

import (
	"fmt"
	"github.com/gnharishkumar13/go-grpc--learn/greet/greetpb"
	"google.golang.org/grpc"
	"log"
)

func main(){
	fmt.Println("Hello, this is client")

	cc, err := grpc.Dial("localhost:50051",grpc.WithInsecure())
	defer cc.Close()

	if err != nil {
		log.Fatalf("client failed to connect to server: %v", err)
	}

	c := greetpb.NewGreetServiceClient(cc)
	fmt.Printf("created client %v",c)

}
