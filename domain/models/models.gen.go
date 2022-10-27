// Package models provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package models

import (
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

// Defines values for DiskDevice.
const (
	DiskDeviceCdrom DiskDevice = "cdrom"
	DiskDeviceDisk  DiskDevice = "disk"
)

// Defines values for DiskType.
const (
	Qcow2   DiskType = "qcow2"
	Unknown DiskType = "unknown"
)

// Defines values for StoragePoolStatus.
const (
	Active StoragePoolStatus = "Active"
	Error  StoragePoolStatus = "Error"
)

// ホストのCPU情報
type Cpu struct {
	Counts  int     `json:"counts"`
	Percent float64 `json:"percent"`
}

// Devices defines model for Devices.
type Devices struct {
	Disk []Disk `json:"disk"`
}

// Disk defines model for Disk.
type Disk struct {
	Device DiskDevice `json:"device"`
	Path   string     `json:"path"`
	Type   DiskType   `json:"type"`
}

// DiskDevice defines model for Disk.Device.
type DiskDevice string

// DiskType defines model for Disk.Type.
type DiskType string

// Host defines model for Host.
type Host struct {
	// ホストのCPU情報
	Cpu Cpu `json:"cpu"`

	// ホストのメモリ情報
	Mem          Memory        `json:"mem"`
	StoragePools []StoragePool `json:"storage_pools"`
}

// ホストのメモリ情報
type Memory struct {
	Free        uint64  `json:"free"`
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"used_percent"`
}

// Metadata defines model for Metadata.
type Metadata struct {
	ApiVersion string             `json:"api_version"`
	Id         openapi_types.UUID `json:"id"`
}

// ホストが持つストレージプールの情報
type StoragePool struct {
	Name      string            `json:"name"`
	Path      string            `json:"path"`
	Status    StoragePoolStatus `json:"status"`
	TotalSize uint64            `json:"total_size"`
	UsedSize  uint64            `json:"used_size"`
}

// StoragePoolStatus defines model for StoragePool.Status.
type StoragePoolStatus string

// 仮想マシンを表すモデル
type Vm struct {
	Devices  Devices  `json:"devices"`
	Memory   int      `json:"memory"`
	Metadata Metadata `json:"metadata"`
	Name     string   `json:"name"`
	Vcpu     int      `json:"vcpu"`
}

// GetAllVMsList200Response defines model for GetAllVMsList200Response.
type GetAllVMsList200Response struct {
	Vms *[]Vm `json:"vms,omitempty"`
}

// GetHost200Response defines model for GetHost200Response.
type GetHost200Response = Host

// 仮想マシンを表すモデル
type GetVMByVMId200Response = Vm

// 仮想マシンを表すモデル
type PatchUpdateVMByVMId200Response = Vm

// 仮想マシンを表すモデル
type PostCreateNewVM200Response = Vm

// PostApiV1VmsJSONBody defines parameters for PostApiV1Vms.
type PostApiV1VmsJSONBody struct {
	Memory int    `json:"memory"`
	Name   string `json:"name"`
	Vcpu   int    `json:"vcpu"`
}

// PatchApiV1VmsVmIdJSONBody defines parameters for PatchApiV1VmsVmId.
type PatchApiV1VmsVmIdJSONBody struct {
	Memory *int    `json:"memory,omitempty"`
	Name   *string `json:"name,omitempty"`
	Vcpu   *int    `json:"vcpu,omitempty"`
}

// PostApiV1VmsJSONRequestBody defines body for PostApiV1Vms for application/json ContentType.
type PostApiV1VmsJSONRequestBody PostApiV1VmsJSONBody

// PatchApiV1VmsVmIdJSONRequestBody defines body for PatchApiV1VmsVmId for application/json ContentType.
type PatchApiV1VmsVmIdJSONRequestBody PatchApiV1VmsVmIdJSONBody
