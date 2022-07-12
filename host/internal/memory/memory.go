package memory

import "github.com/shirou/gopsutil/v3/mem"

func GetInfo() Memory {
	memInfo, _ := mem.VirtualMemory()
	memTotal := memInfo.Total
	memUsed := memInfo.Used
	memFree := memInfo.Free
	memUsedPercent := memInfo.UsedPercent

	return Memory{
		Total:       memTotal,
		Used:        memUsed,
		Free:        memFree,
		UsedPercent: memUsedPercent,
	}
}
