package main

import (
	"context"
	"fmt"
	"log"

	orderv1 "buf.build/gen/go/firacloudtech/grpc-echo-benchmark/protocolbuffers/go/order/v1"
	productv1 "buf.build/gen/go/firacloudtech/grpc-echo-benchmark/protocolbuffers/go/product/v1"

	orderGrpc "buf.build/gen/go/firacloudtech/grpc-echo-benchmark/grpc/go/order/v1/orderv1grpc"
	productGrpc "buf.build/gen/go/firacloudtech/grpc-echo-benchmark/grpc/go/product/v1/productv1grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var port = 5002
var gatewayPort = 8080

func main() {

	if err := run(); err != nil {
		log.Fatal(err)
	}

}

type combinedClient struct {
	productClient productGrpc.ProductServiceClient
	orderClient   orderGrpc.OrderServiceClient
}

func run() error {
	connectTo := fmt.Sprintf("localhost:%d", port)

	conn, err := grpc.Dial(connectTo, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return fmt.Errorf("failed to connect to ProductService on %s: %w", connectTo, err)
	}
	defer conn.Close()
	log.Println("Connected to Backend")

	// create a combined client
	combinedClient := &combinedClient{
		productClient: productGrpc.NewProductServiceClient(conn),
		orderClient:   orderGrpc.NewOrderServiceClient(conn),
	}

	if _, err := combinedClient.productClient.CreateProduct(context.Background(), &productv1.CreateRequest{
		Name: "Yogen",
	}); err != nil {
		return fmt.Errorf("failed to CreateProduct: %w", err)
	}
	log.Println("Successfully Created product")

	if _, err := combinedClient.orderClient.CreateOrder(context.Background(), &orderv1.CreateRequest{
		Name: "Yogen",
	}); err != nil {
		return fmt.Errorf("failed to CreateOrder: %w", err)

	}

	log.Println("Successfully Created order")
	return nil
}
