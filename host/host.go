package host

import (
	"github.com/CharVstack/CharV-lib/host/internal/cpu"
	"github.com/CharVstack/CharV-lib/host/internal/memory"
	"github.com/CharVstack/CharV-lib/host/internal/storage"
)

func GetInfo(opt GetInfoOptions) Host {
	cpuInfo := cpu.GetInfo()
	memoryInfo := memory.GetInfo()

	poolConfigPaths := storage.GetPoolFiles(opt.storageDir)

	var storagePools []*storage.PoolInfo
	for _, file := range poolConfigPaths {
		storagePoolInfo := storage.GetPoolInfo(file, opt.storageDir)
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
	}
}

type GetInfoOptions struct {
	storageDir string
}
