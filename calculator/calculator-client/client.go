package main

import (
	"context"
	"fmt"
	"log"

	"github.com/rishavkumar7/go-grpc/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Calculator client started")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %v\n", err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)
	if err != nil {
		log.Fatalf("Error while creating client: %v\n", err)
	}
	fmt.Printf("Client created: %v\n", c)

	doUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &calculatorpb.CalculatorRequest{
		FirstNumber:  11,
		SecondNumber: 22,
	}
	res, err := c.Calculate(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Calculator RPC: %v\n", err)
	}
	log.Printf("Response from Calculator: %v", res.SumResult)
}
