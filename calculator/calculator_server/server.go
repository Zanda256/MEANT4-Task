package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

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

	var fact factorizer = factorial.NewComputer()

	for _, v := range lst {
		if v < 0 {
			return status.Errorf(
				codes.InvalidArgument,
				fmt.Sprintf("Recieved a negative number %+v. Accepts only positive numbers.", v),
			)
		}
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

	lis, err := net.Listen("tcp", "localhost:5100")
	if err != nil {
		fmt.Printf("server failed to listen on tcp port 5100 : %+v\n", err)
	}
	s := grpc.NewServer()
	pb.RegisterFactorialServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	signalChan := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-signalChan
		fmt.Printf("\n%+v", sig)
		done <- true
	}()
	fmt.Println("Recieved stop signal. Exiting gracefully.")
	s.Stop()

}
