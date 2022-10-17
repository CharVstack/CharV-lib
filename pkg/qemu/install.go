package qemu

import (
	"bytes"
	"text/template"
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

func CreateInfoJSON(opts Vm) Vm {
	var createUUID string
	createUUID = "this-is-uuid"

	createJSON := Vm{
		Devices: Devices{
			Disk: make([]Disk, len(opts.Devices.Disk)),
		},
		Memory: opts.Memory,
		Metadata: Metadata{
			ApiVersion: "v1",
			Id:         createUUID,
		},
		Name: opts.Name,
		Vcpu: opts.Vcpu,
	}

	for i := 0; i < len(opts.Devices.Disk); i++ {
		createJSON.Devices.Disk[i].Type = opts.Devices.Disk[i].Type
		createJSON.Devices.Disk[i].Path = opts.Devices.Disk[i].Path
	}

	return createJSON
}

func Install(opts InstallOpts) error {
	tmpl, err := template.New("install").Parse(`qemu-system-x86_64 -accel kvm -daemonize -display none -name guest={{.Name}} -smp {{.VCpu}} -m {{.Memory}} -cdrom /var/lib/charVstack/iso/{{.Image}} -boot order=d -drive file=/var/lib/charVstack/images/{{.Disk}}.qcow2,format=qcow2 -drive file=/var/lib/charVstack/bios/bios.bin,format=raw,if=pflash,readonly=on`)
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
