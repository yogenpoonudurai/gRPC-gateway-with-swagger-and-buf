package main

import (
	"flag"
	"net/http"

	pb "github.com/firacloudtech/grpc-echo-benchmark/proto/product"
)

var ( port = flag.Int("port", 50051, "Server port"))

type server struct {
	pb.UnimplementedGreeterServer
}


func NewServer() *server{
	return &server{}
}

func (s *server) SayHello(req *http.Request) error{

}