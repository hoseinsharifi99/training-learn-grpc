package main

import (
	"context"
	"fmt"
	"log"

	pb "grpc/greet_pb"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cant connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	doUnary(c)

	//fmt.Println("conectc", c)
}

func doUnary(c pb.GreetServiceClient) {
	fmt.Println("do unary")
	req := &pb.GreetRequest{
		Greating: &pb.Greeting{
			FirstName: "hossein",
			LastName:  "sharifi",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error to calling greet rpc: %v", err)
	}

	log.Printf("response: %v", res.Result)
}
