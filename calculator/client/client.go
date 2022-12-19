package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "grpc/calculatorpb"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cant connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)

	doUnary(c)
	doServerStreaming(c)
	//fmt.Println("conectc", c)
}

func doUnary(c pb.CalculatorServiceClient) {
	fmt.Println("do unary")
	req := &pb.SumRequest{
		FirstNumber:  10,
		SecondNumber: 20,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error to calling greet rpc: %v", err)
	}

	log.Printf("response: %v", res.SumResult)
}

func doServerStreaming(c pb.CalculatorServiceClient) {
	fmt.Println("Starting to do a PrimeDecomposition Server Streaming RPC...")
	req := &pb.PrimeNumberDecompositionRequest{
		Number: 12390392840,
	}
	stream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling PrimeDecomposition RPC: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happened: %v", err)
		}
		fmt.Println(res.GetPrimeFactor())
	}
}
