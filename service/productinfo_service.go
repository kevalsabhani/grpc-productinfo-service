package main

import (
	"context"

	"github.com/google/uuid"
	pb "github.com/kevalsabhani/productinfo/service/ecommerce"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ProductInfoServer use to implement ecommerce/product_info
type server struct {
	products map[string]*pb.Product
	pb.UnimplementedProductInfoServer
}

// AddProduct implements ecommerce.AddProduct
func (s *server) AddProduct(ctx context.Context, in *pb.Product) (*pb.ProductID, error) {
	out, err := uuid.NewV7()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error occurred while generating Product ID: %s", err.Error())
	}

	in.Id = out.String()

	if s.products == nil {
		s.products = make(map[string]*pb.Product)
	}
	s.products[in.Id] = in

	return &pb.ProductID{Value: in.Id}, status.New(codes.OK, "").Err()
}

// GetProduct implements ecommerce.GetProduct
func (s *server) GetProduct(ctx context.Context, in *pb.ProductID) (*pb.Product, error) {
	product, exists := s.products[in.Value]
	if exists {
		return product, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "Product does not exist: %s", in.Value)
}
