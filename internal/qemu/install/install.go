package install

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/CharVstack/CharV-lib/domain"
	"github.com/CharVstack/CharV-lib/domain/models"

	"github.com/google/uuid"
)

func CreateInfoJSON(opts domain.InstallOpts) (models.Vm, error) {
	uuidInt, err := uuid.NewRandom()
	if err != nil {
		return models.Vm{}, err
	}

	uuidString := uuidInt.String()

	vmInfo := models.Vm{
		Devices: models.Devices{
			Disk: []models.Disk{
				{
					Type: "qcow2",
					Path: "/var/lib/charVstack/image/" + opts.Disk,
				},
			},
		},
		Memory: opts.Memory,
		Metadata: models.Metadata{
			ApiVersion: "v1",
			Id:         uuidString,
		},
		Name: opts.Name,
		Vcpu: opts.VCpu,
	}

	var MarshalJSON []byte
	MarshalJSON, err = json.Marshal(vmInfo)
	if err != nil {
		return models.Vm{}, err
	}

	createJSONPath := "/var/lib/charVstack/machines/"

	fileName := createJSONPath + vmInfo.Name + "-" + vmInfo.Metadata.Id + ".json"

	var createFile *os.File
	createFile, err = os.Create(fileName)
	if err != nil {
		return models.Vm{}, err
	}
	defer func() {
		err = createFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	_, err = createFile.Write(MarshalJSON)
	if err != nil {
		return models.Vm{}, err
	}

	return vmInfo, err
}
