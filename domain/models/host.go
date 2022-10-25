package models

type Host struct {
	Cpu          Cpu
	Memory       Memory
	StoragePools []*StoragePool
}

type GetInfoOptions struct {
	StorageDir string
}
