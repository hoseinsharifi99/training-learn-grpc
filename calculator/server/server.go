package main

import (
	"context"
	"fmt"

	pb "grpc/calculatorpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type UserManagementServer struct {
	pb.UnimplementedCalculatorServiceServer
}

func (*UserManagementServer) Sum(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	fmt.Printf("Received Sum RPC: %v\n", req)
	firstNumber := req.FirstNumber
	secondNumber := req.SecondNumber
	sum := firstNumber + secondNumber
	res := &pb.SumResponse{
		SumResult: sum,
	}
	return res, nil
}
func (*UserManagementServer) PrimeNumberDecomposition(req *pb.PrimeNumberDecompositionRequest, stream pb.CalculatorService_PrimeNumberDecompositionServer) error {
	fmt.Printf("Received PrimeNumberDecomposition RPC: %v\n", req)

	number := req.GetNumber()
	divisor := int64(2)

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&pb.PrimeNumberDecompositionResponse{
				PrimeFactor: divisor,
			})
			number = number / divisor
		} else {
			divisor++
			fmt.Printf("Divisor has increased to %v\n", divisor)
		}
	}
	return nil

}
func main() {
	fmt.Println("server is up")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &UserManagementServer{})
	//pb.RegisterGreetServiceServer(s, &UserManagementServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serv: %v", err)
	}

}
