package cpu

import "github.com/shirou/gopsutil/v3/cpu"

func GetInfo() (Cpu, error) {
	cpuCounts, err := cpu.Counts(true)
	if err != nil {
		return Cpu{}, err
	}
	cpuPercent, err := cpu.Percent(0, false)
	if err != nil {
		return Cpu{}, err
	}

	return Cpu{
		Counts:  cpuCounts,
		Percent: cpuPercent[0],
	}, nil
}
