package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	helloV1 "rpcPro/gen/proto/hello/v1"
)

func main() {
	conn, e := grpc.NewClient("localhost:8888",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	fmt.Println(e)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	client := helloV1.NewGreeterClient(conn)
	resp, err := client.SayHello(context.Background(), &helloV1.HelloRequest{
		Name: "gRPC",
	})
	if err != nil {
		panic(err)
	}

	println(resp.GetMessage())

}
