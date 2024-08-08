package main

import (
	"context"
	"fmt"
	helloV1 "github.com/Echin-h/grpc-template-proto/gen/proto/hello/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
)

type server struct {
	helloV1.UnimplementedGreeterServiceServer
}

var _ helloV1.GreeterServiceServer = (*server)(nil)

//func NewServer() *server {
//	return &server{}
//}

func (s *server) SayHello(_ context.Context, in *helloV1.SayHelloRequest) (*helloV1.SayHelloResponse, error) {
	return &helloV1.SayHelloResponse{Message: in.Name + " world"}, nil
}

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	helloV1.RegisterGreeterServiceServer(s, &server{})

	go func() {
		fmt.Println("Serving grpc on http://localhost:8080")
		log.Fatalln(s.Serve(lis))
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.NewClient(
		"0.0.0.0:8080", // target service address
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	// gateway uses mux-multiplexer to config the gwServer
	gwmux := runtime.NewServeMux()
	err = helloV1.RegisterGreeterServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	go func() {
		log.Fatalln(gwServer.ListenAndServe())
	}()
}

// at last , should run two endpoints: - http://localhost:8080 - http://localhost:8090
// register rpc-message to grpc-gateway and serve the http server on 8090
// grpc-client dial grpc-server on 8080
// if curl http://localhost:8090 grpc-gateway will transfer the request to grpc-server by conn which created by grpc.NewClient
