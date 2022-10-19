package install

type Disk struct {
	Type string `json:"type"`
	Path string `json:"path"`
}

type Devices struct {
	Disk []*Disk `json:"disk"`
}

type Metadata struct {
	ApiVersion string `json:"api_version"`
	Id         string `json:"id"`
}

type Vm struct {
	Devices  Devices  `json:"devices"`
	Memory   int      `json:"memory"`
	Metadata Metadata `json:"metadata"`
	Name     string   `json:"name"`
	VCpu     int      `json:"vcpu"`
}

type RequestOpts struct {
	Name   string
	Memory int
	VCpu   int
	Image  string
	Disk   string
}
