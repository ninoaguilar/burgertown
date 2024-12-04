package main

import (
	"context"
	"log"

	pb "github.com/ninoaguilar/burgertown/common/api"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
}

func NewGRPCHandler(grpcServer *grpc.Server) {
	handler := &grpcHandler{}
	pb.RegisterOrderServiceServer(grpcServer, handler)
}

func (h *grpcHandler) CreateOrder(context.Context, *pb.CreateOrderRequest) (*pb.Order, error) {
	log.Println("New order received!")
	o := &pb.Order{
		Id: "42",
	}

	return o, nil
}
