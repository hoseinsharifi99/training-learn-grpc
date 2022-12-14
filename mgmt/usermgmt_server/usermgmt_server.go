package main

import (
	"context"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"os"

	pb "grpc/usermgmt"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

const (
	port = ":50051"
)

func NewUserManagementServer() *UserManagementServer {
	return &UserManagementServer{}
}

type UserManagementServer struct {
	pb.UnimplementedUserManagmentServer
}

func (server *UserManagementServer) Run() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserManagmentServer(s, server)
	log.Printf("server listening at %v", lis.Addr())
	return s.Serve(lis)
}

func (server *UserManagementServer) CreateNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("Received: %v", in.GetName())
	readByte, err := ioutil.ReadFile("users.json")
	var users_list *pb.UsersList = &pb.UsersList{}
	var user_id = int32(rand.Intn(100))
	created_user := &pb.User{Name: in.GetName(), Age: in.GetAge(), Id: user_id}
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("file not found. create a new file ")
			users_list.Users = append(users_list.Users, created_user)
			jsonBytes, err := protojson.Marshal(users_list)
			if err != nil {
				log.Fatalf("json marshaling failed: %v", err)
			}
			if err := ioutil.WriteFile("users.json", jsonBytes, 0664); err != nil {
				log.Fatal("failed write to file: %v", err)
			}
			return created_user, nil
		} else {
			log.Fatalln("error reading file: ", err)
		}
	}
	if err := protojson.Unmarshal(readByte, users_list); err != nil {
		log.Fatalf("failed to parse userlisr: %v", err)
	}
	users_list.Users = append(users_list.Users, created_user)
	jsonBytes, err := protojson.Marshal(users_list)
	if err != nil {
		log.Fatalf("json marshaling failed: %v", err)
	}
	if err := ioutil.WriteFile("users.json", jsonBytes, 0664); err != nil {
		log.Fatal("failed write to file: %v", err)
	}

	return created_user, nil
}

func (server *UserManagementServer) GetUsers(ctx context.Context, in *pb.GetUsersParams) (*pb.UsersList, error) {
	jsonByte, err := ioutil.ReadFile("users.json")
	if err != nil {
		log.Fatalf("failed read file")
	}
	var users_list *pb.UsersList = &pb.UsersList{}
	if err := protojson.Unmarshal(jsonByte, users_list); err != nil {
		log.Fatal("unmarshaling failed: %v", err)
	}

	return users_list, nil
}

func main() {
	var user_mgmt_server *UserManagementServer = NewUserManagementServer()
	if err := user_mgmt_server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
