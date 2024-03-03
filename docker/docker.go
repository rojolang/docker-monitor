package docker

import (
    "context"
    "fmt"

    "github.com/docker/docker/api/types"
    "github.com/docker/docker/client"
)

// ContainerStat holds container statistics.
type ContainerStat struct {
    ID          string
    Image       string
    Status      string
    CPUUsage    string
    MemoryUsage string
}

// FetchStats retrieves statistics for all running containers.
func FetchStats(ctx context.Context) ([]ContainerStat, error) {
    cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
    if err != nil {
        return nil, fmt.Errorf("failed to create Docker client: %w", err)
    }

    containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
    if err != nil {
        return nil, fmt.Errorf("failed to list containers: %w", err)
    }

    var stats []ContainerStat
    // Placeholder: Iterate over containers and fetch stats.
    // Simulate or implement real fetching.

    return stats, nil
}
