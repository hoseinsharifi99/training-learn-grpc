package main

import (
	"context"
	pb "grpc/usermgmt"
	"log"
	"math/rand"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type UserManagmentServer struct {
	pb.UnimplementedUserManagmentServer
	user_List *pb.UsersList
}

func NewUserManagementServer() *UserManagmentServer {
	return &UserManagmentServer{
		user_List: &pb.UsersList{},
	}
}

func (server *UserManagmentServer) CreateNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("Received: %v", in.GetName())
	var user_id = int32(rand.Intn(100))
	created_user := &pb.User{Name: in.GetName(), Age: in.GetAge(), Id: user_id}
	server.user_List.Users = append(server.user_List.Users, created_user)
	return created_user, nil
}

func (server *UserManagmentServer) Run() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserManagmentServer(s, server)
	log.Printf("server listening at %v", lis.Addr())
	return s.Serve(lis)
}

func (s *UserManagmentServer) GetUsers(ctx context.Context, in *pb.GetUsersParams) (*pb.UsersList, error) {
	return s.user_List, nil
}

func main() {
	var user_mgmt_server *UserManagmentServer = NewUserManagementServer()
	if err := user_mgmt_server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
