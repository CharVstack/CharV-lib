package qemu

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"text/template"

	"github.com/google/uuid"
)

func CreateDisk(opts string) error {
	tmpl, err := template.New("create").Parse(`qemu-img create -f qcow2 /var/lib/charVstack/images/{{.}}.qcow2 16G`)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, opts); err != nil {
		return err
	}
	cmd := buf.String()

	return run(cmd)
}

func CreateInfoJSON(opts InstallOpts) (Vm, error) {
	uuidInt, err := uuid.NewRandom()
	if err != nil {
		return Vm{}, err
	}

	uuidString := uuidInt.String()

	vmInfo := Vm{
		Devices: Devices{
			Disk: []*Disk{
				{
					Type: "qcow2",
					Path: "/var/lib/charVstack/image/" + opts.Disk,
				},
			},
		},
		Memory: opts.Memory,
		Metadata: Metadata{
			ApiVersion: "v1",
			Id:         uuidString,
		},
		Name: opts.Name,
		VCpu: opts.VCpu,
	}

	var MarshalJSON []byte
	MarshalJSON, err = json.Marshal(vmInfo)
	if err != nil {
		return Vm{}, err
	}

	createJSONPath := "/var/lib/charVstack/machines/"

	fileName := createJSONPath + vmInfo.Name + "-" + vmInfo.Metadata.Id + ".json"

	var createFile *os.File
	createFile, err = os.Create(fileName)
	if err != nil {
		return Vm{}, err
	}
	defer func() {
		err = createFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	_, err = createFile.Write(MarshalJSON)
	if err != nil {
		return Vm{}, err
	}

	return vmInfo, err
}

func Install(opts InstallOpts) (Vm, error) {
	tmpl, err := template.New("install").Parse(`qemu-system-x86_64 -accel kvm -daemonize -display none -name guest={{.Name}} -smp {{.VCpu}} -m {{.Memory}} -cdrom /var/lib/charVstack/iso/{{.Image}} -boot order=d -drive file=/var/lib/charVstack/images/{{.Disk}}.qcow2,format=qcow2 -drive file=/var/lib/charVstack/bios/bios.bin,format=raw,if=pflash,readonly=on`)
	if err != nil {
		return Vm{}, err
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, opts); err != nil {
		return Vm{}, err
	}
	cmd := buf.String()

	var getJSON Vm
	getJSON, err = CreateInfoJSON(opts)
	if err != nil {
		return Vm{}, err
	}

	return getJSON, run(cmd)
}
