package host

import (
	"github.com/CharVstack/CharV-lib/internal/host/cpu"
	"github.com/CharVstack/CharV-lib/internal/host/memory"
	storage2 "github.com/CharVstack/CharV-lib/internal/host/storage"
)

func GetInfo() Host {
	cpuInfo := cpu.GetInfo()
	memoryInfo := memory.GetInfo()

	poolConfigPaths := storage2.GetPoolFiles()

	var storagePools []*storage2.PoolInfo
	for _, file := range poolConfigPaths {
		storagePoolInfo := storage2.GetPoolInfo(file)
		isExists, _ := storage2.IsPoolExists(storagePoolInfo.Path)
		if isExists {
			storagePoolInfo.Status = "Active"
			storagePoolInfo.TotalSize, storagePoolInfo.UsedSize = storage2.GetSize(storagePoolInfo.Path)
		} else {
			storagePoolInfo.Status = "Error"
		}

		storagePools = append(storagePools, storagePoolInfo)
	}

	return Host{
		Cpu:          cpuInfo,
		Memory:       memoryInfo,
		StoragePools: storagePools,
	}
}
