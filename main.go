package main

import (
	"context"
	helloV1 "github.com/Echin-h/grpc-template/gen/proto/hello/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
)

type Server struct {
	helloV1.UnimplementedGreeterServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) SayHello(ctx context.Context, req *helloV1.HelloRequest) (*helloV1.HelloReply, error) {
	return &helloV1.HelloReply{Message: "Hello " + req.Name}, nil
}

func (s *Server) SayHelloAgain(ctx context.Context, req *helloV1.HelloRequest) (*helloV1.HelloReply, error) {
	return &helloV1.HelloReply{Message: "Hello again " + req.Name}, nil
}

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	helloV1.RegisterGreeterServer(s, &Server{})

	log.Println("Server started on port 8080")
	go func() {
		log.Fatalln(s.Serve(l))
	}()

	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	gwmux := runtime.NewServeMux()
	// register client
	err = helloV1.RegisterGreeterHandler(context.Background(), gwmux, conn)
	if err != nil {
		panic(err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	log.Println("Gateway started on port 8090")
	log.Fatalln(gwServer.ListenAndServe())

}
