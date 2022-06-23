package cpu

import "github.com/shirou/gopsutil/v3/cpu"

func GetInfo() int {
	cpuInfo, _ := cpu.Counts(true)
	return cpuInfo
}
