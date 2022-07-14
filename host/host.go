package host

import (
	"github.com/CharVstack/CharV-lib/host/internal/cpu"
	"github.com/CharVstack/CharV-lib/host/internal/memory"
	"github.com/CharVstack/CharV-lib/host/internal/storage"
)

func GetInfo() (Host, error) {
	cpuInfo, err := cpu.GetInfo()
	if err != nil {
		return Host{}, err
	}
	memoryInfo, err := memory.GetInfo()
	if err != nil {
		return Host{}, err
	}

	poolConfigPaths := storage.GetPoolFiles()

	var storagePools []*storage.PoolInfo
	for _, file := range poolConfigPaths {
		storagePoolInfo := storage.GetPoolInfo(file)
		isExists, _ := storage.IsPoolExists(storagePoolInfo.Path)
		if isExists {
			storagePoolInfo.Status = "Active"
			storagePoolInfo.TotalSize, storagePoolInfo.UsedSize = storage.GetSize(storagePoolInfo.Path)
		} else {
			storagePoolInfo.Status = "Error"
		}

		storagePools = append(storagePools, storagePoolInfo)
	}

	return Host{
		Cpu:          cpuInfo,
		Memory:       memoryInfo,
		StoragePools: storagePools,
	}, nil
}
