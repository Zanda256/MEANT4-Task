package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/Zanda256/MEANT4-Task/calculator/calc_proto"
	"github.com/Zanda256/MEANT4-Task/factorial"
)

type server struct {
	pb.UnimplementedFactorialServer
}

type factorizer interface {
	Compute(int64) string
}

func (s *server) Calculate(req *pb.CalculateRequest, stream pb.Factorial_CalculateServer) error {
	lst := req.GetNumbers()
	var fact factorizer = factorial.NewComputer(5)
	for _, v := range lst {
		str := fact.Compute(v)
		stream.Send(&pb.CalculateResult{
			InputNumber:     v,
			FactorialResult: str,
		})
	}
	return nil
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
