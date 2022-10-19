package qemu

import "github.com/CharVstack/CharV-lib/internal/qemu/install"

type Machine struct {
	Name    string          `json:"name"`
	Memory  int             `json:"memory"`
	VCpu    int             `json:"vcpu"`
	Devices install.Devices `json:"devices"`
}

type InstallOpts struct {
	RequestData install.RequestOpts
}

type StartOpts struct {
	Disk string
}
