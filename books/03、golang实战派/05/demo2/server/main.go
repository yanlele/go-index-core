package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "practice/05/demo2"
)

type ProgrammerServiceServer struct {
}

func (p *ProgrammerServiceServer) GetProgrammerInfo(ctx context.Context, req *pb.Request) (resp *pb.Response, err error) {
	name := req.Name
	if name == "yanle" {
		resp = &pb.Response{
			Uid:      6,
			Username: name,
			Job:      "programmer",
			GoodAt:   []string{"js", "go"},
		}
	}

	err = nil
	return
}

func main() {
	port := ":8078"
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("listen error : %v\n", err)
	}
	fmt.Printf("listen %s\n", port)
	s := grpc.NewServer()
	pb.RegisterProgrammerServiceServer(s, &ProgrammerServiceServer{})
	s.Serve(listen)
}
