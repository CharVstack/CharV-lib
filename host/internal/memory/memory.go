package memory

import "github.com/shirou/gopsutil/v3/mem"

func GetInfo() (Memory, error) {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return Memory{}, err
	}
	memTotal := memInfo.Total
	memUsed := memInfo.Used
	memFree := memInfo.Free
	memUsedPercent := memInfo.UsedPercent

	return Memory{
		Total:       memTotal,
		Used:        memUsed,
		Free:        memFree,
		UsedPercent: memUsedPercent,
	}, nil
}
