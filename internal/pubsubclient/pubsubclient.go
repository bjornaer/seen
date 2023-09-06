package pubsubclient

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync/atomic"
	"time"

	"cloud.google.com/go/pubsub"
)

type SeenMsg struct {
	JobID   string              `json:"jobId"`
	MsgType string              `json:"type"`
	Shots   []map[string]string `json:"shots"`
}

func DecodeMessage(data []byte) (*SeenMsg, error) {
	var seenMsg SeenMsg
	err := json.Unmarshal(data, &seenMsg)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}
	return &seenMsg, nil
}

func PullMsgs(client *pubsub.Client, subID string) error {
	ctx := context.Background()

	sub := client.Subscription(subID)
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var received int32
	err := sub.Receive(ctx, func(_ context.Context, msg *pubsub.Message) {
		seenMsg, err := DecodeMessage(msg.Data)
		if err != nil {
			log.Printf("Failed to decode message: %v", err)
			return
		}
		log.Printf("Got message: %v", seenMsg)
		atomic.AddInt32(&received, 1)
		msg.Ack()
	})

	if err != nil {
		return fmt.Errorf("sub.Receive: %w", err)
	}
	log.Printf("Received %d messages", received)

	return nil
}
