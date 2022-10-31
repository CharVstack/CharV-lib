package qemu

import (
	"encoding/json"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/CharVstack/CharV-lib/domain/models"
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

func ConvertToStruct(directoryPath string) ([]models.Vm, error) {
	var resJSONList []models.Vm
	dir, err := os.ReadDir(directoryPath)
	if err != nil {
		return []models.Vm{}, err
	}

	var openFile *os.File
	var read = make([]byte, 1024)
	var cont int
	var resJSON models.Vm

	for _, file := range dir {
		openFile, err = os.Open(directoryPath + file.Name())
		if err != nil {
			return []models.Vm{}, err
		}

		cont, err = openFile.Read(read)
		if err != nil {
			return []models.Vm{}, err
		}

		err = json.Unmarshal(read[:cont], &resJSON)
		if err != nil {
			return []models.Vm{}, err
		}

		resJSONList = append(resJSONList, resJSON)

	}

	err = openFile.Close()
	if err != nil {
		return []models.Vm{}, err
	}

	return resJSONList, err
}
