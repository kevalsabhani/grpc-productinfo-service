package main

import (
	"context"
	"log"
	"time"

	pb "github.com/kevalsabhani/productinfo/client/ecommerce"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.NewClient(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewProductInfoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	product := &pb.Product{
		Name:        "Apple iphone 11",
		Description: "Meet Apple iPhone 11. Ultra Wide and Night mode.",
		Price:       float32(60000.0),
	}
	resp, err := c.AddProduct(ctx, product)
	if err != nil {
		log.Fatalf("Could not add product: %v", err)
	}
	log.Printf("Product ID: %s added successfully", resp.Value)

	r, err := c.GetProduct(ctx, &pb.ProductID{Value: resp.Value})
	if err != nil {
		log.Fatalf("could not get product: %v", err)
	}
	log.Printf("Product: %s", r.String())
}
