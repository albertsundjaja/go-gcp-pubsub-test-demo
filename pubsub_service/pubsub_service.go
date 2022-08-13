package pubsub_service

import (
	"context"
	"fmt"
	"strconv"

	"cloud.google.com/go/pubsub"
	"github.com/albertsundjaja/go-gcp-pubsub-test-demo/pubsub_client"
)

type PubsubService struct {
	client pubsub_client.IPubsubClient
}

func (p *PubsubService) Publish(topicId string) error {
	// do some calculation
	x := 1 + 2

	// publish the message
	ctx := context.Background()
	t := p.client.Topic(topicId)
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(strconv.Itoa(x)),
	})

	// check the message published successfully and get the msg id
	id, err := result.Get(ctx)
	if err != nil {
		fmt.Printf("Failed to publish: %v \n", err)
		return err
	}
	fmt.Printf("Published message. msg ID: %v \n", id)

	return nil
}

func NewPubsubService(client pubsub_client.IPubsubClient) *PubsubService {
	return &PubsubService{
		client: client,
	}
}
