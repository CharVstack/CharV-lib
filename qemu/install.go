package qemu

import (
	"bytes"
	"text/template"
)

func CreateDisk(opts string) error {
	tmpl, err := template.New("create").Parse(`qemu-img create -f qcow2 /var/lib/charVstack/images/{{.}}.qcow2 8G`)
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

func Install(opts InstallOpts) error {
	tmpl, err := template.New("install").Parse(`qemu-system-x86_64 -smp {{.VCpu}} -m {{.Memory}} -cdrom "{{.Name}}" -boot order=d -drive file=/var/lib/charVstack/images/{{.Disk}}.qcow2,format=qcow2`)
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
