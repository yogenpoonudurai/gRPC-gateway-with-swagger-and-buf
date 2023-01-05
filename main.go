package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github/firacloudtech/grpc-echo-benchmark/proto/product"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "Server port")
)

type server struct {
	pb.UnimplementedProductServiceServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) SayHello(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {

	return &pb.CreateResponse{Message: req.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterProductServiceServer(s, &server{})

	log.Printf("Serving gPRC on port %d", port)
	log.Fatal(s.Serve(lis))
}
