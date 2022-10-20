package install

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/CharVstack/CharV-lib/domain"

	"github.com/google/uuid"
)

func CreateInfoJSON(opts domain.InstallOpts) (domain.Vm, error) {
	uuidInt, err := uuid.NewRandom()
	if err != nil {
		return domain.Vm{}, err
	}

	uuidString := uuidInt.String()

	vmInfo := domain.Vm{
		Devices: domain.Devices{
			Disk: []*domain.Disk{
				{
					Type: "qcow2",
					Path: "/var/lib/charVstack/image/" + opts.Disk,
				},
			},
		},
		Memory: opts.Memory,
		Metadata: domain.Metadata{
			ApiVersion: "v1",
			Id:         uuidString,
		},
		Name: opts.Name,
		VCpu: opts.VCpu,
	}

	var MarshalJSON []byte
	MarshalJSON, err = json.Marshal(vmInfo)
	if err != nil {
		return domain.Vm{}, err
	}

	createJSONPath := "/var/lib/charVstack/machines/"

	fileName := createJSONPath + vmInfo.Name + "-" + vmInfo.Metadata.Id + ".json"

	var createFile *os.File
	createFile, err = os.Create(fileName)
	if err != nil {
		return domain.Vm{}, err
	}
	defer func() {
		err = createFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	_, err = createFile.Write(MarshalJSON)
	if err != nil {
		return domain.Vm{}, err
	}

	return vmInfo, err
}
