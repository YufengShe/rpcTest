package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"test/pb/helloService"
)

func main() {
	grpcServer := grpc.NewServer()
	helloService.RegisterHelloServiceServer(grpcServer, new(HelloServer))
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err.Error())
	}
	_ = grpcServer.Serve(listen)
}
