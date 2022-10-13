package cpu

import "github.com/shirou/gopsutil/v3/cpu"

func GetInfo() Cpu {
	cpuCounts, _ := cpu.Counts(true)
	cpuPercent, _ := cpu.Percent(0, false)

	return Cpu{
		Counts:  cpuCounts,
		Percent: cpuPercent[0],
	}
}
