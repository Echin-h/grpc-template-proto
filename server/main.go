package main

import (
	"context"
	"fmt"
	helloV1 "github.com/Echin-h/grpc-temmplate/gen/proto/hello/v1"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	helloV1.UnimplementedGreeterServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) SayHello(ctx context.Context, req *helloV1.HelloRequest) (res *helloV1.HelloReply, err error) {
	fmt.Println(req.GetName())
	//return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
	return &helloV1.HelloReply{
		Message: "Hello, " + req.GetName(),
	}, nil
}

func (s *Server) SayHelloAgain(ctx context.Context, req *helloV1.HelloRequest) (res *helloV1.HelloReply, err error) {
	fmt.Println(req.GetName())
	//return nil, status.Errorf(codes.Unimplemented, "method SayHelloAgain not implemented")
	return &helloV1.HelloReply{
		Message: "Hello again, " + req.GetName(),
	}, nil
}

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}

	// create a grpc server object
	s := grpc.NewServer()

	// attach the greeter service to the server
	helloV1.RegisterGreeterServer(s, &Server{})

	// Serve grpc Server
	log.Fatal(s.Serve(l))
}
