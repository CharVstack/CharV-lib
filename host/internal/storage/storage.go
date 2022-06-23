package storage

import (
	"github.com/shirou/gopsutil/v3/disk"
)

func GetSize(path string) (size uint64, usedSize uint64) {
	diskUsage, _ := disk.Usage(path)
	return diskUsage.Total, diskUsage.Used
}
