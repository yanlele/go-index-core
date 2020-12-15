package main

import (
	"context"
	"fmt"
	pb "gin-example/proto"
	"google.golang.org/grpc"
	"io"
	"log"
)

const PORT = "9002"

func main() {
	conn, err := grpc.Dial(
		fmt.Sprintf(":%s", PORT),
		grpc.WithInsecure(),
	)

	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}

	defer conn.Close()

	client := pb.NewStreamServiceClient(conn)

	err = printLists(
		client,
		&pb.StreamRequest{
			Pt: &pb.StreamPoint{
				Name:  "gRPC Stream Client: List",
				Value: 2020,
			},
		},
	)

	err = printRecord(
		client,
		&pb.StreamRequest{
			Pt: &pb.StreamPoint{
				Name:  "gRPC Stream Client: Record",
				Value: 2020,
			},
		},
	)

	err = printRecord(
		client,
		&pb.StreamRequest{
			Pt: &pb.StreamPoint{
				Name:  "gRPC Stream Client: Route",
				Value: 2020,
			},
		},
	)

	if err != nil {
		log.Fatalf("printRoute.err %v", err)
	}
}

func printLists(client pb.StreamServiceClient, r *pb.StreamRequest) error {
	stream, err := client.List(context.Background(), r)
	if err != nil {
		return err
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Printf("resp: pt.name: %s, pt.value: %d", resp.Pt.Name, resp.Pt.Value)
	}
	return nil
}

func printRecord(client pb.StreamServiceClient, r *pb.StreamRequest) error {
	stream, err := client.Record(context.Background())
	if err != nil {
		return err
	}

	for n := 0; n < 6; n++ {
		err := stream.Send(r)
		if err != nil {
			return err
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}

	log.Printf(
		"resp: pt.name: %s, pt.value: %d",
		resp.Pt.Name,
		resp.Pt.Value,
	)

	return nil
}

func pintRoute(client pb.StreamServiceClient, r *pb.StreamRequest) error {
	stream, err := client.Route(context.Background())
	if err != nil {
		return err
	}

	for n := 0; n < 6; n++ {
		err = stream.Send(r)
		if err != nil {
			return err
		}
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Printf(
			"resp: pt.name: %s, pt.value: %d",
			resp.Pt.Name,
			resp.Pt.Value,
		)
	}

	stream.CloseSend()

	return nil
}
