package domain

import "github.com/CharVstack/CharV-lib/domain/models"

type Machine struct {
	Name    string         `json:"name"`
	Memory  int            `json:"memory"`
	VCpu    int            `json:"vcpu"`
	Devices models.Devices `json:"devices"`
}

type InstallOpts struct {
	Name   string
	Memory int
	VCpu   int
	Image  string
	Disk   string
}

type StartOpts struct {
	Disk string
}
