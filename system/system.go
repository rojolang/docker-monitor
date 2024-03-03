package system

// DockerSystemStats represents overall Docker system resource usage.
type DockerSystemStats struct {
    TotalMemoryUsage string
    TotalCPUUsage    string
}

// FetchSystemStats simulates the fetching of Docker system stats.
func FetchSystemStats() DockerSystemStats {
    // Placeholder: Implement real fetching or simulate.
    return DockerSystemStats{
        TotalMemoryUsage: "8GB",
        TotalCPUUsage:    "40%",
    }
}
