package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/Dre4m6904/assignment4/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) AddUser(ctx context.Context, in *pb.User) (*pb.User, error) {
	fmt.Printf("Adding user: %v\n", in)
	return in, nil
}

func (s *server) GetUser(ctx context.Context, in *pb.UserID) (*pb.User, error) {
	fmt.Printf("Getting user with ID: %v\n", in.Id)
	return &pb.User{
		Id:    in.Id,
		Name:  "John Doe",
		Email: "john@example.com",
	}, nil
}

func (s *server) ListUsers(req *pb.Empty, stream pb.UserService_ListUsersServer) error {
	users := []*pb.User{
		{Id: 1, Name: "User1", Email: "user1@example.com"},
		{Id: 2, Name: "User2", Email: "user2@example.com"},
	}

	for _, user := range users {
		if err := stream.Send(user); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	fmt.Println("Server started on port :4040")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
