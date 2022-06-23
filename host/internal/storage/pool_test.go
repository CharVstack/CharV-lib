package storage

import (
	"reflect"
	"testing"
)

func TestGetPoolFiles(t *testing.T) {
	t.Setenv("STORAGE_DIR", "../../../storage")
	files := []string{"default.json"}
	if !reflect.DeepEqual(GetPoolFiles(), files) {
		t.Error(`miss`)
	}
}

func TestGetPoolInfo(t *testing.T) {
	t.Setenv("STORAGE_DIR", "../../../storage")
	storagePool := PoolInfo{
		Name: "default",
		Path: "/var/lib/libvirt/default",
	}
	storagePoolInfo := GetPoolInfo("default.json")
	if !reflect.DeepEqual(*storagePoolInfo, storagePool) {
		t.Error(`miss`)
	}
}
