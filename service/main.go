package main

import (
	"log"
	"net"

	pb "github.com/kevalsabhani/productinfo/service/ecommerce"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterProductInfoServer(grpcServer, &server{})
	log.Printf("Starting grpc sever on port %s", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
