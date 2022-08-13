package main

import (
	"context"
	"fmt"

	pubsub "github.com/albertsundjaja/go-gcp-pubsub-test-demo/pubsub_client"
)

const (
	PROJECT_ID = "my-project-id"
	TOPIC_ID   = "my-topic"
)

func main() {
	ctx := context.Background()
	client, err := pubsub_client.NewPubsubClient(ctx, PROJECT_ID)
	if err != nil {
		fmt.Printf("Error creating pubsub client: %v", err)
		return
	}

	pubsubService := pubsub.NewPubsubService(client)
	pubsubService.Publish()
}
