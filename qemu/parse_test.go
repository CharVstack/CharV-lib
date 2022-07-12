package qemu

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	testMachine := Machine{
		Name:   "ubuntu",
		Memory: 1024,
		VCpu:   1,
		Devices: Devices{
			Disk: []*Disk{
				{
					Type:   "file",
					Device: "disk",
					Path:   "/path/to/ubuntu.qcow2",
				},
			},
		},
	}
	machine, err := parse("../test/resources/machines/ubuntu.json")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(testMachine, machine) {
		t.Fail()
	}
}
