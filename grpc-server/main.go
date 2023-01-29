package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	// import module using buf
	orderv1 "github.com/firacloudtech/grpc-echo-benchmark/gen/go/order/v1"
	productv1 "github.com/firacloudtech/grpc-echo-benchmark/gen/go/product/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
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

	conn, err := grpc.DialContext(context.Background(), "127.0.0.1:5002", grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	defer conn.Close()

	gwmux := runtime.NewServeMux()
	err = productv1.RegisterProductServiceHandler(context.Background(), gwmux, conn)

	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%v", gatewayPort),
		Handler: gwmux,
	}

	log.Printf("Serving gRPC-gateway on http://127.0.0.1%v", gwServer.Addr)

	err = gwServer.ListenAndServe()
	if err != nil {
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

func (s *combinedServer) CreateOrder(ctx context.Context, req *orderv1.CreateOrderRequest) (*orderv1.CreateOrderResponse, error) {

	name := req.GetName()

	log.Printf("Got a request to create a order: %v\n", name)

	return &orderv1.CreateOrderResponse{}, nil

}
