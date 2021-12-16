package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	"test/pb/helloService"
	"time"
)

/*func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close()

	client := helloService.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &helloService.String{Value: "YF"})
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("reply: ", reply)
}*/

func main()  {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close()

	client := helloService.NewHelloServiceClient(conn)
	channel, err := client.Channel(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}
	go func() {
		for {
			err := channel.Send(&helloService.String{Value: "YF"})
			if err != nil {
				log.Fatal(err.Error())
			}
			time.Sleep(time.Second)
		}
	}()

	for {
		reply, err := channel.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err.Error())
		}
		log.Println(reply.GetValue())
	}
}