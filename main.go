package main

import (
	"context"
	"fmt"

	"github.com/albertsundjaja/go-gcp-pubsub-test-demo/pubsub_client"
	"github.com/albertsundjaja/go-gcp-pubsub-test-demo/pubsub_service"
)

const (
	PROJECT_ID = "my-project-id"
	TOPIC_ID   = "my-topic"
)

func main() {
	ctx := context.Background()
	client, err := pubsub_client.NewPubsubClient(ctx, PROJECT_ID)
	if err != nil {
		fmt.Printf("Error creating pubsub client: %v \n", err)
		return
	}

	pubsubService := pubsub_service.NewPubsubService(client)
	pubsubService.Publish(TOPIC_ID)
}
