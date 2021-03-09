package main

import (
	pb "demo/proto"
	rpcServer "demo/rpc"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":50000"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	server, _ := rpcServer.New()
	pb.RegisterUsersServer(grpcServer, server)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
