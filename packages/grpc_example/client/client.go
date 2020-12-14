package main

import (
	"context"
	"fmt"
	"gin-example/proto"
	"google.golang.org/grpc"
	"log"
)

const PORT = "9001"

func main() {
	conn, err := grpc.Dial(fmt.Sprintf(":%s", PORT), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()
	client := proto.NewSearchServiceClient(conn)
	response, err := client.Search(context.Background(), &proto.SearchRequest{Request: "gRPC"})
	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}
	log.Printf("res: %s", response.GetResponse())
}
