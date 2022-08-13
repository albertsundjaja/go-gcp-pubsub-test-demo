package pubsub

import (
	"context"
	"fmt"
	"strconv"

	"cloud.google.com/go/pubsub"
)

// PublishResult
type IPubsubPublishResult interface {
	Get(ctx context.Context) (msgID string, err error)
}

type PubsubPublishResult struct {
	*pubsub.PublishResult
}

func (p *PubsubPublishResult) Get(ctx context.Context) (msgID string, err error) {
	id, err := p.PublishResult.Get(ctx)
	if err != nil {
		fmt.Printf("Failed to publish: %v", err)
		return "", err
	}
	return id, nil
}

// PubsubTopic
type IPubsubTopic interface {
	Publish(ctx context.Context, msg *pubsub.Message) IPubsubPublishResult
}

type PubsubTopic struct {
	*pubsub.Topic
}

func (p *PubsubTopic) Publish(ctx context.Context, msg *pubsub.Message) IPubsubPublishResult {
	result := p.Topic.Publish(ctx, msg)
	return &PubsubPublishResult{
		PublishResult: result,
	}
}

// PubsubClient
type IPubsubClient interface {
	Topic(topicId string) IPubsubTopic
}

type PubsubClient struct {
	client *pubsub.Client
}

func (p *PubsubClient) Topic(topicId string) IPubsubTopic {
	t := p.client.Topic(topicId)
	return &PubsubTopic{
		Topic: t,
	}
}

func NewPubsubClient(ctx context.Context, projectID string) (*PubsubClient, error) {
	// create new pubsub client
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		fmt.Printf("pubsub client error: %v", err)
		return nil, err
	}
	return &PubsubClient{
		client: client,
	}, nil
}

type PubsubService struct {
	client IPubsubClient
}

func (p *PubsubService) Publish(topicId string) {
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
		fmt.Printf("Failed to publish: %v", err)
		return
	}
	fmt.Printf("Published message. msg ID: %v", id)
}

func NewPubsubService(client IPubsubClient) *PubsubService {
	return &PubsubService{
		client: client,
	}
}
