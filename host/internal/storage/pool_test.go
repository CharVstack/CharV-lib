package storage

import (
	"reflect"
	"testing"
)

func TestGetStoragePoolFiles(t *testing.T) {
	files := []string{"default.json"}
	if !reflect.DeepEqual(GetPoolFiles(), files) {
		t.Error(`miss`)
	}
}

func TestGetStoragePoolInfo(t *testing.T) {
	storagePool := PoolInfo{
		Name:      "default",
		TotalSize: 20000,
		Path:      "/var/cpu/libvirt/default",
	}
	storagePoolInfo := GetPoolInfo("default.json")
	if !reflect.DeepEqual(storagePoolInfo, storagePool) {
		t.Error(`miss`)
	}
}
