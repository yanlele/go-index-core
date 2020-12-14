package main

import (
	"context"
	"gin-example/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type SearchService struct{}

func (s *SearchService) Search(ctx context.Context, r *proto.SearchRequest) (*proto.SearchResponse, error) {
	return &proto.SearchResponse{Response: r.GetRequest() + " Server"}, nil
}

const PORT = "9001"

// 一个最简单的demo
func main() {
	server := grpc.NewServer()
	proto.RegisterSearchServiceServer(server, &SearchService{})

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	_ = server.Serve(lis)
}
