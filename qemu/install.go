package qemu

import (
	"bytes"
	"text/template"
)

func install(opts InstallOpts) error {
	tmpl, err := template.New("install").Parse(`qemu-system-x86_64 -smp {{.VCpu}} -m {{.Memory}} -cdrom "{{.Name}}" -boot order=d -drive file={{.Disk}},format=raw`)
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
