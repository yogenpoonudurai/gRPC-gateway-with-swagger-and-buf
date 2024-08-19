package main

import (
	"context"
	"fmt"
	"log"

	orderv1 "github.com/firacloudtech/grpc-gateway-swagger-buf/gen/go/order/v1"
	productv1 "github.com/firacloudtech/grpc-gateway-swagger-buf/gen/go/product/v1"

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
	productClient productv1.ProductServiceClient
	orderClient   orderv1.OrderServiceClient
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
		productClient: productv1.NewProductServiceClient(conn),
		orderClient:   orderv1.NewOrderServiceClient(conn),
	}

	if _, err := combinedClient.productClient.CreateProduct(context.Background(), &productv1.CreateProductRequest{
		Name: "Yogen",
	}); err != nil {
		return fmt.Errorf("failed to CreateProduct: %w", err)
	}
	log.Println("Successfully Created product")

	if _, err := combinedClient.orderClient.CreateOrder(context.Background(), &orderv1.CreateOrderRequest{
		Name: "Client1",
	}); err != nil {
		return fmt.Errorf("failed to CreateOrder: %w", err)

	}

	log.Println("Successfully Created order")
	return nil
}
