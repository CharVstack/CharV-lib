package host

import (
	"github.com/CharVstack/CharV-lib/domain"
	"github.com/CharVstack/CharV-lib/domain/models"
	"github.com/CharVstack/CharV-lib/internal/host/cpu"
	"github.com/CharVstack/CharV-lib/internal/host/memory"
	"github.com/CharVstack/CharV-lib/internal/host/storage"
)

func GetInfo(opt domain.GetInfoOptions) (domain.Host, error) {
	cpuInfo, err := cpu.GetInfo()
	if err != nil {
		return domain.Host{}, err
	}

	memoryInfo, err := memory.GetInfo()
	if err != nil {
		return domain.Host{}, err
	}

	poolConfigPaths, err := storage.GetPoolFiles(opt.StorageDir)
	if err != nil {
		return domain.Host{}, err
	}

	var storagePools []*models.StoragePool
	for _, file := range poolConfigPaths {
		storagePoolInfo, err := storage.GetPoolInfo(file, opt.StorageDir)
		if err != nil {
			return domain.Host{}, err
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

	return domain.Host{
		Cpu:          cpuInfo,
		Memory:       memoryInfo,
		StoragePools: storagePools,
	}, nil
}
