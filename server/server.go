package main

import (
	"context"
	"io"
	"log"
	"test/pb/helloService"
)

type HelloServer struct {
	helloService.UnimplementedHelloServiceServer
}

func (s *HelloServer) Hello(ctx context.Context,str *helloService.String) (*helloService.String, error) {
	response := helloService.String{Value: "Hello " + str.GetValue()}
	return &response, nil
}

func (s *HelloServer) Channel(channel helloService.HelloService_ChannelServer) error {
	for {
		args, err := channel.Recv()
		if err != nil {
			if err == io.EOF {
				log.Println("Finished")
				return nil
			}
			return err
		}
		reply := helloService.String{Value: "Hello " + args.Value}
		err = channel.Send(&reply)
		if err != nil {
			return err
		}
	}
}
