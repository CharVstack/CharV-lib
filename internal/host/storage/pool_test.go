package storage

import (
	"reflect"
	"testing"
)

func TestGetPoolFiles(t *testing.T) {
	files := []string{"default.json"}
	if !reflect.DeepEqual(GetPoolFiles("../../../test/resources/storage"), files) {
		t.Error(`miss`)
	}
}

func TestGetPoolInfo(t *testing.T) {
	storagePool := PoolInfo{
		Name: "default",
		Path: "/var/lib/libvirt/default",
	}
	storagePoolInfo := GetPoolInfo("default.json", "../../../test/resources/storage")
	if !reflect.DeepEqual(*storagePoolInfo, storagePool) {
		t.Error(`miss`)
	}
}
