package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	// import module using buf

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	orderv1 "github.com/firacloudtech/grpc-echo-benchmark/gen/go/order/v1"
	productv1 "github.com/firacloudtech/grpc-echo-benchmark/gen/go/product/v1"
)

var (
	port        = 5002
	gatewayPort = 3001
	wg          sync.WaitGroup
)

type combinedServer struct {
	productv1.UnimplementedProductServiceServer
	orderv1.UnimplementedOrderServiceServer
}

func NewServer() *combinedServer {
	return &combinedServer{}
}

func main() {
	wg.Add(2)

	go func() {
		if err := run(); err != nil {
			log.Fatal(err)
		}
		wg.Done()
	}()

	// wait till the grpc server is ready

	go func() {

		if err := runGrpcGateway(); err != nil {
			log.Fatal(err)
		}
		wg.Done()
	}()
	wg.Wait()
}

func run() error {
	grpcAddr := fmt.Sprintf("127.0.0.1:%d", port)
	listener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %v", grpcAddr, err)

	}
	defer listener.Close()

	// Register the services
	server := grpc.NewServer()
	productv1.RegisterProductServiceServer(server, &combinedServer{})
	orderv1.RegisterOrderServiceServer(server, &combinedServer{})

	log.Println("Listening on port", grpcAddr)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)

	}

	return nil
}

func runGrpcGateway() error {
	// grpc gateway server

	gwmux := runtime.NewServeMux()
	gwAddr := fmt.Sprintf("127.0.0.1:%d", gatewayPort)

	conn, err := grpc.DialContext(context.Background(), fmt.Sprintf("127.0.0.1:%d", port), grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	err = productv1.RegisterProductServiceHandler(context.Background(), gwmux, conn)

	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    gwAddr,
		Handler: gwmux,
	}

	log.Println("Serving gRPC-gateway on http://localhost:", gatewayPort)

	if err := gwServer.ListenAndServe(); err != nil {
		return fmt.Errorf("failed to serve gRPC gateway server: %w", err)
	}

	return nil
}

// a function that creates a product and return the response
func (s *combinedServer) CreateProduct(ctx context.Context, req *productv1.CreateProductRequest) (*productv1.CreateProductResponse, error) {

	name := req.GetName()

	log.Printf("Got a request to create a product: %v\n", name)

	return &productv1.CreateProductResponse{
		Message: "Success",
	}, nil

}

func (s *combinedServer) CreateOrder(ctx context.Context, req *orderv1.CreateRequest) (*orderv1.CreateResponse, error) {

	name := req.GetName()

	log.Printf("Got a request to create a order: %v\n", name)

	return &orderv1.CreateResponse{}, nil

}
