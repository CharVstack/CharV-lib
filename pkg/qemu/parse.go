package qemu

import (
	"encoding/json"
	"github.com/CharVstack/CharV-lib/domain"
	"io/ioutil"
	"os"
	"path/filepath"
)

func parse(path string) (domain.Machine, error) {
	var machine domain.Machine
	abspath, err := filepath.Abs(path)
	f, err := os.Open(abspath)
	if err != nil {
		return domain.Machine{}, err
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return domain.Machine{}, err
	}
	if err := json.Unmarshal(b, &machine); err != nil {
		return domain.Machine{}, err
	}
	return machine, nil
}
