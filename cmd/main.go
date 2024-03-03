package main

import (
    "context"
    "log"
    "docker-monitor/docker"
    "docker-monitor/system"
    "docker-monitor/ui"
)

func main() {
    ctx := context.Background()

    containerStats, err := docker.FetchStats(ctx)
    if err != nil {
        log.Fatalf("Error fetching Docker stats: %v", err)
    }

    sysStats := system.FetchSystemStats()

    if err := ui.StartUI(containerStats, sysStats); err != nil {
        log.Fatalf("Error starting UI: %v", err)
    }
}
