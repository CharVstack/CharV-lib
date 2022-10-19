package qemu

import (
	"bytes"
	"text/template"

	"github.com/CharVstack/CharV-lib/internal/qemu/install"
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

func Install(opts InstallOpts) (install.Vm, error) {
	requestData := install.RequestOpts{
		Name:   opts.RequestData.Name,
		Memory: opts.RequestData.Memory,
		VCpu:   opts.RequestData.VCpu,
		Image:  opts.RequestData.Image,
		Disk:   opts.RequestData.Disk,
	}

	tmpl, err := template.New("install").Parse(`qemu-system-x86_64 -accel kvm -daemonize -display none -name guest={{.Name}} -smp {{.VCpu}} -m {{.Memory}} -cdrom /var/lib/charVstack/iso/{{.Image}} -boot order=d -drive file=/var/lib/charVstack/images/{{.Disk}}.qcow2,format=qcow2 -drive file=/var/lib/charVstack/bios/bios.bin,format=raw,if=pflash,readonly=on`)
	if err != nil {
		return install.Vm{}, err
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, requestData)
	if err != nil {
		return install.Vm{}, err
	}
	cmd := buf.String()

	var getJSON install.Vm
	getJSON, err = install.CreateInfoJSON(requestData)
	if err != nil {
		return install.Vm{}, err
	}

	return getJSON, run(cmd)
}
