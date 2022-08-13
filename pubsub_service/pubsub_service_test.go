package pubsub_service

import (
	"fmt"
	"testing"

	mock "github.com/albertsundjaja/go-gcp-pubsub-test-demo/mock"
	"github.com/golang/mock/gomock"
)

func TestPublishSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPublishResult := mock.NewMockIPubsubPublishResult(ctrl)
	mockTopic := mock.NewMockIPubsubTopic(ctrl)
	mockClient := mock.NewMockIPubsubClient(ctrl)
	pubsubService := NewPubsubService(mockClient)

	mockClient.EXPECT().Topic(gomock.Any()).Return(mockTopic)
	mockTopic.EXPECT().Publish(gomock.Any(), gomock.Any()).Return(mockPublishResult)
	mockPublishResult.EXPECT().Get(gomock.Any()).Return("id", nil)

	err := pubsubService.Publish("TOPIC")
	if err != nil {
		t.Errorf("Publish should not return an error")
	}
}

func TestPublishFailed(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPublishResult := mock.NewMockIPubsubPublishResult(ctrl)
	mockTopic := mock.NewMockIPubsubTopic(ctrl)
	mockClient := mock.NewMockIPubsubClient(ctrl)
	pubsubService := NewPubsubService(mockClient)

	mockClient.EXPECT().Topic(gomock.Any()).Return(mockTopic)
	mockTopic.EXPECT().Publish(gomock.Any(), gomock.Any()).Return(mockPublishResult)
	mockPublishResult.EXPECT().Get(gomock.Any()).Return("", fmt.Errorf("error occurred"))

	err := pubsubService.Publish("TOPIC")
	if err == nil {
		t.Errorf("Publish should return an error")
	}
}
