package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/Zanda256/MEANT4-Task/calculator/calc_proto"
)

type server struct {
	pb.UnimplementedFactorialServer
}

func main() {
	fmt.Println("factorial server is up.")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		fmt.Printf("server failed to listen on tcp port 50051 : %+v\n", err)
	}
	s := grpc.NewServer()
	pb.RegisterFactorialServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
