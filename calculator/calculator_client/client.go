package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/Zanda256/MEANT4-Task/calculator/calc_proto"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("grpc client here.")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("client failed to dial rpc %+v\n", err)
	}
	facClnt := pb.NewFactorialClient(conn)
	integers := []int64{99, 200, 16, 20, 48, 63, 89, 32, 72, 5}

	streamResults(facClnt, integers)
}

func streamResults(fc pb.FactorialClient, arr []int64) {
	req := &pb.CalculateRequest{
		Numbers: arr,
	}
	resStream, err := fc.Calculate(context.Background(), req)
	if err != nil {
		log.Fatal("error while calling calculate")
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			//end of stream
			break
		}
		if err != nil {
			log.Fatal("error while reading stream %+v", err)
		}
		fmt.Printf("The factorial of %+v is %+v\n", msg.GetInputNumber(), msg.GetFactorialResult())
	}
}
