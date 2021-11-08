package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/Zanda256/MEANT4-Task/calculator/calc_proto"
	"github.com/Zanda256/MEANT4-Task/cli"
	"google.golang.org/grpc"
)

var expectedType = "integers"

func main() {
	fmt.Println("grpc client here.")

	conn, err := grpc.Dial("localhost:5100", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("client failed to dial rpc %+v\n", err)
	}
	defer conn.Close()
	facClnt := pb.NewFactorialClient(conn)

	integers, err := cli.GetUserInput(&expectedType)
	if err != nil {
		log.Fatalf("error: %+v", err)
	}

	streamResults(facClnt, integers)
}

func streamResults(fc pb.FactorialClient, arr []int64) {
	req := &pb.CalculateRequest{
		Numbers: arr,
	}
	resStream, err := fc.Calculate(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling calculate\n%+v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			//end of stream
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream %+v", err)
		}
		if msg.InputNumber > 1000000 {
			fmt.Printf("The factorial of %+v is approximately %+v.\n", msg.GetInputNumber(), msg.GetFactorialResult())
		} else {
			fmt.Printf("The factorial of %+v is %+v.\n", msg.GetInputNumber(), msg.GetFactorialResult())
		}
	}
}
