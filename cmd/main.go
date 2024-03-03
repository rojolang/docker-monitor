package main

import (
	"context"
	"github.com/docker/docker/client"
	"github.com/rojolang/docker-monitor/dockerstats" // Adjust the import path based on your project structure
	log "github.com/sirupsen/logrus"
)

func main() {
	// Initialize Docker client
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Error creating Docker client: %v", err)
	}

	// Create a context
	ctx := context.Background()

	// Use the dockerstats package to fetch and display Docker container stats
	dockerstats.FetchAndDisplayStats(ctx, cli)
}
