package qemu

import (
	"bytes"
	"text/template"

	"github.com/CharVstack/CharV-lib/domain/models"
)

func start(opts models.InstallOpts) error {
	tmpl, err := template.New("start").Parse(`qemu-system-x86_64 {{.Disk}}`)
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
