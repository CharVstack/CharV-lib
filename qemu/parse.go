package qemu

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

func parse(path string) (Machine, error) {
	var machine Machine
	abspath, err := filepath.Abs(path)
	f, err := os.Open(abspath)
	if err != nil {
		return Machine{}, err
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return Machine{}, err
	}
	if err := json.Unmarshal(b, &machine); err != nil {
		return Machine{}, err
	}
	return machine, nil
}
