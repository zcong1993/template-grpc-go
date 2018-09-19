package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"<%= projectPath %>/pb"
	"google.golang.org/grpc"
)

// HelloService is our grpc service
type HelloService struct{}

// Echo impl Echo method
func (service *HelloService) Echo(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	fmt.Printf("recieve message: %+v\n", in)
	return &pb.Message{Value: in.Value}, nil
}

func runGatewayServer(rpcPort, port string) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := pb.RegisterEchoServiceHandlerFromEndpoint(ctx, mux, rpcPort, opts)

	if err != nil {
		log.Fatal("Serve http error:", err)
	}

	http.ListenAndServe(port, mux)
}

func runRpcServer(port string) {
	ss := grpc.NewServer()
	pb.RegisterEchoServiceServer(ss, &HelloService{})

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	if err = ss.Serve(listener); err != nil {
		log.Fatal("ListenTCP error:", err)
	}
}

func main() {
	if os.Getenv("GATEWAY") == "true" {
		println("run gateway server on :9009")
		go runGatewayServer(":1234", ":9009")
	}
	runRpcServer(":1234")
}
