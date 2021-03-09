package main

import (
	"context"
	"log"
	"time"

	pb "demo/proto"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50000"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUsersClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetUsers(ctx, &pb.EmptyReq{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Get users from GRPC server: %v", r.Users)
}
