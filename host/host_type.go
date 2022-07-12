package host

import (
	"github.com/CharVstack/CharV-lib/host/internal/cpu"
	"github.com/CharVstack/CharV-lib/host/internal/memory"
	"github.com/CharVstack/CharV-lib/host/internal/storage"
)

type Host struct {
	Cpu          cpu.Cpu
	Memory       memory.Memory
	StoragePools []*storage.PoolInfo
}
