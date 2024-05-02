package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/Dre4m6904/assignment4/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	user := &pb.User{
		Id:    1,
		Name:  "Alice",
		Email: "alice@example.com",
	}

	addedUser, err := client.AddUser(context.Background(), user)
	if err != nil {
		log.Fatalf("failed to add user: %v", err)
	}
	fmt.Printf("Added user: %v\n", addedUser)

	getUserResponse, err := client.GetUser(context.Background(), &pb.UserID{Id: 1})
	if err != nil {
		log.Fatalf("failed to get user: %v", err)
	}
	fmt.Printf("Get user response: %v\n", getUserResponse)

	listUsersStream, err := client.ListUsers(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("failed to list users: %v", err)
	}
	for {
		user, err := listUsersStream.Recv()
		if err != nil {
			break
		}
		fmt.Printf("List user: %v\n", user)
	}
}
