package domain

import (
	"github.com/CharVstack/CharV-lib/internal/host/cpu"
	"github.com/CharVstack/CharV-lib/internal/host/memory"
	"github.com/CharVstack/CharV-lib/internal/host/storage"
)

type Host struct {
	Cpu          cpu.Cpu
	Memory       memory.Memory
	StoragePools []*storage.PoolInfo
}
