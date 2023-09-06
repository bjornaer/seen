package pubsubclient_test

import (
	"context"
	"testing"

	"cloud.google.com/go/pubsub"
	"cloud.google.com/go/pubsub/pstest"
	"github.com/stretchr/testify/assert"
	"google.golang.org/api/option"
	"google.golang.org/grpc"

	"github.com/bjornaer/seen/internal/pubsubclient"
)

func TestDecodeMessage(t *testing.T) {
	rawMessage := `{ "jobId": "8b54ba12-4e6b-4dad-ad18-36a2d0362cc6", "type":
	"encoding", "shots": [ { "source": "name.mov" }, { "source":
	"place.mov" }, { "source": "animal.mov" } ] }`
	expected := &pubsubclient.SeenMsg{
		JobID:   "8b54ba12-4e6b-4dad-ad18-36a2d0362cc6",
		MsgType: "encoding",
		Shots:   []map[string]string{{"source": "name.mov"}, {"source": "place.mov"}, {"source": "animal.mov"}},
	}
	data := []byte(rawMessage)
	msg, err := pubsubclient.DecodeMessage(data)
	assert.Nil(t, err)
	assert.Equal(t, expected, msg)
}

func TestPullMsgs(t *testing.T) {
	// Start a fake server
	ctx := context.Background()
	server := pstest.NewServer()
	defer server.Close()
	conn, err := grpc.Dial(server.Addr, grpc.WithInsecure())
	assert.Nil(t, err)

	// Create a new client
	client, err := pubsub.NewClient(ctx, "seen-test", option.WithGRPCConn(conn))
	assert.Nil(t, err)

	// Create a new topic and subscription
	topic, err := client.CreateTopic(ctx, "seen-test-topic")
	assert.Nil(t, err)
	_, err = client.CreateSubscription(ctx, "seen-test-subscription", pubsub.SubscriptionConfig{
		Topic: topic,
	})
	assert.Nil(t, err)

	// Publish a message to the topic
	topic.Publish(ctx, &pubsub.Message{
		Data: []byte(`{"jobID": "job123", "msgType": "typeABC", "shots": [{"key": "value"}]}`),
	})

	// Run the PullMsgs function
	err = pubsubclient.PullMsgs(client, "seen-test-subscription")
	assert.Nil(t, err)
}
