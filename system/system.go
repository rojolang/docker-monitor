package system

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

type DockerSystemStats struct {
	TotalMemoryUsage uint64
	TotalCPUUsage    float64
	DiskUsage        *disk.UsageStat
}

func FetchSystemStats() (DockerSystemStats, error) {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return DockerSystemStats{}, fmt.Errorf("error fetching memory stats: %v", err)
	}

	cpuPercent, err := cpu.Percent(0, false)
	if err != nil {
		return DockerSystemStats{}, fmt.Errorf("error fetching CPU stats: %v", err)
	}

	diskStat, err := disk.Usage("/")
	if err != nil {
		return DockerSystemStats{}, fmt.Errorf("error fetching disk stats: %v", err)
	}

	return DockerSystemStats{
		TotalMemoryUsage: vmStat.Used,
		TotalCPUUsage:    cpuPercent[0], // Assuming single CPU or average.
		DiskUsage:        diskStat,
	}, nil
}
