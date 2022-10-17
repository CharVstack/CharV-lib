package host

import (
	"github.com/CharVstack/CharV-lib/internal/host/cpu"
	"github.com/CharVstack/CharV-lib/internal/host/memory"
	"github.com/CharVstack/CharV-lib/internal/host/storage"
)

func GetInfo(opt GetInfoOptions) (Host, error) {
	cpuInfo, err := cpu.GetInfo()
	if err != nil {
		return Host{}, err
	}

	memoryInfo, err := memory.GetInfo()
	if err != nil {
		return Host{}, err
	}

	poolConfigPaths, err := storage.GetPoolFiles(opt.storageDir)
	if err != nil {
		return Host{}, err
	}

	var storagePools []*storage.PoolInfo
	for _, file := range poolConfigPaths {
		storagePoolInfo, err := storage.GetPoolInfo(file, opt.storageDir)
		if err != nil {
			return Host{}, err
		}

		isExists := storage.IsPoolExists(storagePoolInfo.Path)

		if isExists {
			storagePoolInfo.Status = "Active"
			storagePoolInfo.TotalSize, storagePoolInfo.UsedSize, err = storage.GetSize(storagePoolInfo.Path)
			if err != nil {
				storagePoolInfo.Status = "Error"
			}
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

type GetInfoOptions struct {
	storageDir string
}
