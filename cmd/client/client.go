package main

import (
	"context"
	"log"

	"github.com/GuilhermeBiavati/grpc-go/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect to gRPC server: %v", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)

	AddUser(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Joao",
		Email: "fasdfa",
	}

	res, err := client.AddUser(context.Background(), req)

	if err != nil {
		log.Fatalf("could not make  gRPC server: %v", err)
	}

	log.Println(res)
}
