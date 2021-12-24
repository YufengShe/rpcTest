package main

import (
	"fmt"
	"github.com/docker/docker/pkg/pubsub"
	"strings"
	"time"
)

func main() {
	p := pubsub.NewPublisher(100*time.Millisecond, 10)

	golang := p.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, "golang:") {
				return true
			}
		}
		return false
	})

	java := p.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, "java:") {
				return true
			}
		}
		return false
	})

	go p.Publish("hi")
	go p.Publish("golang: https://golang.org")
	go p.Publish("java:https://java.org")
	time.Sleep(1*time.Second)

	go func() {
		fmt.Println("golang topic:", <-golang )
	}()

	go func() {
		fmt.Println("java topic:", <-java)
	}()

	time.Sleep(10*time.Second)
}