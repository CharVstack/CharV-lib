package storage

type status string

const (
	Active = status("Active")
	Error  = status("Error")
)

type PoolInfo struct {
	Name      string `json:"name" description:"プールの名前"`
	TotalSize uint64 `json:"size" description:"プールのサイズ(byte)"`
	UsedSize  uint64 `json:"used_size" description:"使用済みサイズ(byte)"`
	Path      string `json:"path" description:"ストレージプールのパス"`
	Status    status `json:"status" description:"プールの状態"`
}
