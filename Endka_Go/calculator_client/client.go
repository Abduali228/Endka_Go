package main

import (
	"context"
	"fmt"
	"github.com/deepakvashist/grpc-go/calculator/calculatorpb"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Calculator client!")
	cc, err := grpc.Dial("localhost:4000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	// doServerStreaming(c)
	doClientStreaming(c)

}


func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a PrimeDecomposition")
	req := &calculatorpb.PrimeNumberDecompositionRequest{
		Number: 12,
	}
	stream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling PrimeDecomposition streaming RPC: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happened: %v", err)
		}
		fmt.Println(res.GetPrimeFactor())
	}
}

func doClientStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a ComputeAverage")
	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("Error while opening stream: %v", err)
	}

	numbers := []int32{2, 2, 2, 3, 5}

	for _, number := range numbers {
		fmt.Printf("Sending number: %v\n", number)
		_ = stream.Send(&calculatorpb.ComputeAverageRequest{
			Number: number,
		})
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response: %v", err)
	}

	fmt.Printf("The Average is: %v\n", res.GetAverage())
}

