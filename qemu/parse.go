package qemu

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

func parse(path string) (machine Machine, err error) {
	abspath, err := filepath.Abs(path)
	f, err := os.Open(abspath)
	if err != nil {
		return Machine{}, err
	}
	defer func() {
		err = f.Close()
	}()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return Machine{}, err
	}
	if err := json.Unmarshal(b, &machine); err != nil {
		return Machine{}, err
	}
	return
}
