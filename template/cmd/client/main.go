package main

import (
	"context"
	"fmt"
	"log"

	"<%= projectPath %>/pb"
	"google.golang.org/grpc"
)

func main() {
	c, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer c.Close()

	client := pb.NewEchoServiceClient(c)

	resp, err := client.Echo(context.Background(), &pb.Message{Value: "hello grpc"})

	if err != nil {
		fmt.Printf("err: %+v\n", err)
		return
	}

	fmt.Println(resp.Value)
}
