package main

import (
	helloV1 "github.com/Echin-h/grpc-template-proto/gen/proto/hello/v1"
	"google.golang.org/grpc"
	"log"
	"net"
)

type HelloService struct {
	helloV1.UnimplementedGreeterServiceServer
}

var _ helloV1.GreeterServiceServer = (*HelloService)(nil)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer(grpc.MaxRecvMsgSize(1024 * 1024 * 1024))
	helloV1.RegisterGreeterServiceServer(&grpc.Server{}, &HelloService{})
	log.Fatal(grpcServer.Serve(lis))
}
