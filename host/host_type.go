package host

import "github.com/CharVstack/CharV-lib/host/internal/storage"

type Host struct {
	Cpu          int
	Memory       uint64
	StoragePools []*storage.PoolInfo
}
