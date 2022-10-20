package qemu

import (
	"github.com/CharVstack/CharV-lib/domain"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	testMachine := domain.Machine{
		Name:   "ubuntu",
		Memory: 1024,
		VCpu:   1,
		Devices: domain.Devices{
			Disk: []*domain.Disk{
				{
					Type: "file",
					Path: "/path/to/ubuntu.qcow2",
				},
			},
		},
	}
	machine, err := parse("../../test/resources/machines/ubuntu.json")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(testMachine, machine) {
		t.Fail()
	}
}
