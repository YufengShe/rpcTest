package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"test/pb/pubSubService"
	"time"
)

func main() {
	grpcServer := grpc.NewServer()
	pubSubService.RegisterPubSubServiceServer(grpcServer, NewPubSubServer(10*time.Millisecond, 10))
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err.Error())
	}
	_ = grpcServer.Serve(listen)
}
