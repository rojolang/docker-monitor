package dockerstats

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/rojolang/docker-monitor/system" // Adjust this import path to where your system package is located
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

// calculateCPUUsage simplifies the CPU usage calculation for a Docker container.
func calculateCPUUsage(stats types.StatsJSON) float64 {
	cpuDelta := float64(stats.CPUStats.CPUUsage.TotalUsage - stats.PreCPUStats.CPUUsage.TotalUsage)
	systemDelta := float64(stats.CPUStats.SystemUsage - stats.PreCPUStats.SystemUsage)
	if systemDelta > 0 && cpuDelta > 0 {
		return (cpuDelta / systemDelta) * 100.0
	}
	return 0.0
}

// FetchAndDisplayStats integrates system stats with Docker container stats fetching and displaying.
func FetchAndDisplayStats(ctx context.Context, cli *client.Client) {
	sysStats, err := system.FetchSystemStats()
	if err != nil {
		log.Fatalf("Error fetching system stats: %v", err)
	}

	fmt.Printf("System Stats:\nCPU Usage: %.2f%%\nMemory Usage: %d bytes\nDisk Usage: %d bytes used of %d bytes (%.2f%% used)\n",
		sysStats.TotalCPUUsage,
		sysStats.TotalMemoryUsage,
		sysStats.DiskUsage.Used, sysStats.DiskUsage.Total, sysStats.DiskUsage.UsedPercent)

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		log.Fatalf("Error listing containers: %v", err)
	}

	fmt.Println("Docker Container Stats:")
	for _, container := range containers {
		stats, err := cli.ContainerStatsOneShot(ctx, container.ID)
		if err != nil {
			log.Errorf("Failed to fetch stats for container %s: %v", container.ID, err)
			continue
		}

		bodyBytes, err := ioutil.ReadAll(stats.Body)
		if err != nil {
			log.Errorf("Failed to read stats for container %s: %v", container.ID, err)
			continue
		}
		defer stats.Body.Close()

		var statsJSON types.StatsJSON
		if err := json.Unmarshal(bodyBytes, &statsJSON); err != nil {
			log.Errorf("Failed to unmarshal stats for container %s: %v", container.ID, err)
			continue
		}

		cpuUsagePercentage := calculateCPUUsage(statsJSON) / sysStats.TotalCPUUsage * 100
		memoryUsagePercentage := float64(statsJSON.MemoryStats.Usage) / float64(sysStats.TotalMemoryUsage) * 100

		fmt.Printf("%s\t%s\tCPU: %.2f%%\tMemory: %.2f%%\n",
			container.ID[:12],
			container.Image,
			cpuUsagePercentage,
			memoryUsagePercentage)
	}
}
