package host

import (
	"reflect"
	"testing"

	"github.com/CharVstack/ChaV-lib/host/internal/storage"
)

func TestGetPoolFiles(t *testing.T) {
	files := []string{"default.json"}
	if !reflect.DeepEqual(storage.GetPoolFiles(), files) {
		t.Error(`miss`)
	}
}

func TestGetPoolInfo(t *testing.T) {
	storagePool := storage.PoolInfo{
		Name: "default",
		Path: "/var/cpu/libvirt/default",
	}
	storagePoolInfo := storage.GetPoolInfo("default.json")
	if !reflect.DeepEqual(storagePoolInfo, storagePool) {
		t.Error(`miss`)
	}
}
