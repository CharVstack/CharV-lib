package qemu

import (
	"os/exec"
	"regexp"
	"strings"
)

func GetRunningList() ([]string, error) {
	out, err := exec.Command("ps", "-o", "cmd", "axh").Output()
	if err != nil {
		return nil, err
	}
	sout := string(out)
	processes := strings.Split(sout, "\n")
	var guests []string
	for _, v := range processes {
		executable := strings.Split(v, " ")[0]
		r := regexp.MustCompile(`qemu-system-x86_64`)
		if r.MatchString(executable) {
			guests = append(guests, executable)
		}
	}
	return guests, nil
}
