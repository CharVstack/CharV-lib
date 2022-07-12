package memory

type Memory struct {
	Total       uint64  `json:"total" description:"メモリサイズ"`
	Used        uint64  `json:"used" description:"使用済みサイズ"`
	Free        uint64  `json:"free" description:"空きサイズ"`
	UsedPercent float64 `json:"used_percent" description:"使用率 (％)"`
}
