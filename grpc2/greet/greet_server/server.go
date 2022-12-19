package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "grpc/greet_pb"

	"google.golang.org/grpc"
)

type UserManagementServer struct {
	pb.UnimplementedGreetServiceServer
}

func (*UserManagementServer) Greet(c context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)
	firstName := req.GetGreating().GetFirstName()
	result := "Hello " + firstName
	res := &pb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("server is up")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterGreetServiceServer(s, &UserManagementServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serv: %v", err)
	}

}
