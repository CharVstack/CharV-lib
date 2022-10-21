package domain

import (
	"github.com/CharVstack/CharV-lib/domain/models"
)

type Host struct {
	Cpu          models.Cpu
	Memory       models.Memory
	StoragePools []*models.StoragePool
}

type GetInfoOptions struct {
	StorageDir string
}
