package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	productv1 "github.com/firacloudtech/grpc-echo-benchmark/gen/proto/go/product/v1"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 5001, "Server port")
)

type server struct {
	productv1.UnimplementedProductServiceServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) SayHello(ctx context.Context, req *productv1.CreateRequest) (*productv1.CreateResponse, error) {

	return &productv1.CreateResponse{Message: req.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	productv1.RegisterProductServiceServer(s, &server{})

	log.Printf("Serving gPRC on port %d", port)
	log.Fatal(s.Serve(lis))
}
