package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	pb "practice/05/demo2"
)

func main() {
	conn, err := grpc.Dial(":8078", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("dail error %v\n", err)
	}

	defer conn.Close()

	client := pb.NewProgrammerServiceClient(conn)
	req := new(pb.Request)
	req.Name = "yanle"

	resp, err := client.GetProgrammerInfo(context.Background(), req)
	if err != nil {
		log.Fatalf("resp error: %v\n", err)
	}
	fmt.Printf("received %v\n", resp)
}
