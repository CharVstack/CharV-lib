package memory

import "github.com/shirou/gopsutil/v3/mem"

func GetInfo() uint64 {
	memoryInfo, _ := mem.VirtualMemory()
	return memoryInfo.Total
}
