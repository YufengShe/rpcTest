package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"test/pb/pubSubService"
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

	var publishText string
	pubCli := pubSubService.NewPubSubServiceClient(conn)
	for {
		_, _ = fmt.Scanf("%s", &publishText)
		if publishText == "q" {
			log.Println("quit!")
			break
		}
		rst, err := pubCli.Publish(context.Background(), &pubSubService.String{Value: publishText})
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Println(rst.GetValue())
	}

}