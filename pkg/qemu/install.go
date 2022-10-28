package qemu

import (
	"bytes"
	"text/template"

	"github.com/CharVstack/CharV-lib/domain/models"

	"github.com/CharVstack/CharV-lib/internal/qemu"
)

func CreateDisk(name string) (string, error) {
	name = "/var/lib/charVstack/images/" + name + "." + "qcow2"
	tmpl, err := template.New("create").Parse(`qemu-img create -f qcow2 {{.}} 16G`)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, name)
	if err != nil {
		return "", err
	}
	cmd := buf.String()

	return name, run(cmd)
}

func Install(opts models.InstallOpts, filePath string) (models.Vm, error) {
	tmpl, err := template.New("install").Parse(`qemu-system-x86_64 -accel kvm -daemonize -display none -name guest={{.Name}} -smp {{.VCpu}} -m {{.Memory}}G -cdrom /var/lib/charVstack/iso/{{.Image}} -boot order=d -drive file=/var/lib/charVstack/images/{{.Disk}}.qcow2,format=qcow2 -drive file=/var/lib/charVstack/bios/bios.bin,format=raw,if=pflash,readonly=on`)
	if err != nil {
		return models.Vm{}, err
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, opts)
	if err != nil {
		return models.Vm{}, err
	}
	cmd := buf.String()

	var resJSON models.Vm
	resJSON, err = qemu.CreateInfoJSON(opts, filePath)
	if err != nil {
		return models.Vm{}, err
	}

	return resJSON, run(cmd)
}
