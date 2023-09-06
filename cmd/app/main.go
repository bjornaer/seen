package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
	"github.com/bjornaer/seen/internal/pubsubclient"
)

func run() error {
	subID := os.Getenv("SUB_ID")         // e.g., "jobs-failed-subscription"
	projectID := os.Getenv("PROJECT_ID") // e.g., "seen-local"

	if subID == "" || projectID == "" {
		return fmt.Errorf("environment variables SUB_ID or PROJECT_ID not set")
	}

	// Initialize Pub/Sub client
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)

	if err != nil {
		return fmt.Errorf("failed to initialize client: %w", err)
	}

	// Resource clean-up
	defer func() {
		if closeErr := client.Close(); closeErr != nil {
			log.Printf("failed to close client: %v", closeErr)
		}
	}()

	// Pull Messages
	err = pubsubclient.PullMsgs(client, subID)
	if err != nil {
		return fmt.Errorf("failed to pull messages: %w", err)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatalf("An error occurred: %v", err)
	}
}
