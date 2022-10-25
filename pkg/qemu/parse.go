package qemu

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/CharVstack/CharV-lib/domain/models"
)

func parse(path string) (models.Machine, error) {
	var machine models.Machine
	abspath, err := filepath.Abs(path)
	f, err := os.Open(abspath)
	if err != nil {
		return models.Machine{}, err
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return models.Machine{}, err
	}
	if err := json.Unmarshal(b, &machine); err != nil {
		return models.Machine{}, err
	}
	return machine, nil
}
