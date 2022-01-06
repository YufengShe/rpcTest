package main

import (
	"context"
	"github.com/docker/docker/pkg/pubsub"
	"strings"
	"test/pb/pubSubService"
	"time"
)

type PubSubServer struct {
	pubSubService.UnimplementedPubSubServiceServer
	pubSub *pubsub.Publisher
}

func NewPubSubServer(timeout time.Duration, buff int) *PubSubServer {
	return &PubSubServer{
		pubSub:  pubsub.NewPublisher(timeout, buff),
	}
}

func (s *PubSubServer) Publish(ctx context.Context, text *pubSubService.String) (*pubSubService.String, error) {
	s.pubSub.Publish(text.GetValue())
	return &pubSubService.String{Value: "Success"}, nil
}

func (s *PubSubServer) Subscribe(topic *pubSubService.String, stream pubSubService.PubSubService_SubscribeServer) error {
	ch := s.pubSub.SubscribeTopic(func (v interface{}) bool {
		if text, ok := v.(string); ok {
			if strings.HasPrefix(text, topic.GetValue()) {
				return true
			}
		}
		return false
	})
	for text := range ch {
		if err := stream.Send(&pubSubService.String{Value: text.(string)}); err != nil {
			return err
		}
	}
	return nil
}
