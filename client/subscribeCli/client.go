package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"test/pb/pubSubService"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close()
	var subTopic string
	subCli := pubSubService.NewPubSubServiceClient(conn)
	//subscribe topic
	_, _ = fmt.Scanf("%s", &subTopic)
	stream ,err := subCli.Subscribe(context.Background(), &pubSubService.String{Value: subTopic})
	if err != nil {
		log.Fatal(err.Error())
	}
	for {
		text, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				log.Println("Finish")
				break
			} else {
				log.Fatal(err.Error())
			}
		}
		log.Println(subTopic + ": " + text.GetValue())
	}
}

