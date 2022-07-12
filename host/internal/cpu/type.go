package cpu

type Cpu struct {
	Counts  int     `json:"counts" description:"論理コア数"`
	Percent float64 `json:"percent" description:"CPU使用率 (％)"`
}
