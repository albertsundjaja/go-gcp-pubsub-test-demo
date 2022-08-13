package pubsub

import (
	"testing"

	mock "github.com/albertsundjaja/go-gcp-pubsub-test-demo/mock"
	"github.com/golang/mock/gomock"
)

func TestPublish(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockClient := mock.NewMockIPubsubClient(ctrl)
	pubsubService := NewPubsubService(mockClient)

	pubsubService.Publish("TOPIC")

}
