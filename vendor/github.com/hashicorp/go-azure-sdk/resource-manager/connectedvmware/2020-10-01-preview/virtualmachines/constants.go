package virtualmachines

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiskMode string

const (
	DiskModeIndependentNonpersistent DiskMode = "independent_nonpersistent"
	DiskModeIndependentPersistent    DiskMode = "independent_persistent"
	DiskModePersistent               DiskMode = "persistent"
)

func PossibleValuesForDiskMode() []string {
	return []string{
		string(DiskModeIndependentNonpersistent),
		string(DiskModeIndependentPersistent),
		string(DiskModePersistent),
	}
}

func parseDiskMode(input string) (*DiskMode, error) {
	vals := map[string]DiskMode{
		"independent_nonpersistent": DiskModeIndependentNonpersistent,
		"independent_persistent":    DiskModeIndependentPersistent,
		"persistent":                DiskModePersistent,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DiskMode(input)
	return &out, nil
}

type DiskType string

const (
	DiskTypeFlat        DiskType = "flat"
	DiskTypePmem        DiskType = "pmem"
	DiskTypeRawphysical DiskType = "rawphysical"
	DiskTypeRawvirtual  DiskType = "rawvirtual"
	DiskTypeSesparse    DiskType = "sesparse"
	DiskTypeSparse      DiskType = "sparse"
	DiskTypeUnknown     DiskType = "unknown"
)

func PossibleValuesForDiskType() []string {
	return []string{
		string(DiskTypeFlat),
		string(DiskTypePmem),
		string(DiskTypeRawphysical),
		string(DiskTypeRawvirtual),
		string(DiskTypeSesparse),
		string(DiskTypeSparse),
		string(DiskTypeUnknown),
	}
}

func parseDiskType(input string) (*DiskType, error) {
	vals := map[string]DiskType{
		"flat":        DiskTypeFlat,
		"pmem":        DiskTypePmem,
		"rawphysical": DiskTypeRawphysical,
		"rawvirtual":  DiskTypeRawvirtual,
		"sesparse":    DiskTypeSesparse,
		"sparse":      DiskTypeSparse,
		"unknown":     DiskTypeUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DiskType(input)
	return &out, nil
}

type FirmwareType string

const (
	FirmwareTypeBios FirmwareType = "bios"
	FirmwareTypeEfi  FirmwareType = "efi"
)

func PossibleValuesForFirmwareType() []string {
	return []string{
		string(FirmwareTypeBios),
		string(FirmwareTypeEfi),
	}
}

func parseFirmwareType(input string) (*FirmwareType, error) {
	vals := map[string]FirmwareType{
		"bios": FirmwareTypeBios,
		"efi":  FirmwareTypeEfi,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FirmwareType(input)
	return &out, nil
}

type IPAddressAllocationMethod string

const (
	IPAddressAllocationMethodDynamic   IPAddressAllocationMethod = "dynamic"
	IPAddressAllocationMethodLinklayer IPAddressAllocationMethod = "linklayer"
	IPAddressAllocationMethodOther     IPAddressAllocationMethod = "other"
	IPAddressAllocationMethodRandom    IPAddressAllocationMethod = "random"
	IPAddressAllocationMethodStatic    IPAddressAllocationMethod = "static"
	IPAddressAllocationMethodUnset     IPAddressAllocationMethod = "unset"
)

func PossibleValuesForIPAddressAllocationMethod() []string {
	return []string{
		string(IPAddressAllocationMethodDynamic),
		string(IPAddressAllocationMethodLinklayer),
		string(IPAddressAllocationMethodOther),
		string(IPAddressAllocationMethodRandom),
		string(IPAddressAllocationMethodStatic),
		string(IPAddressAllocationMethodUnset),
	}
}

func parseIPAddressAllocationMethod(input string) (*IPAddressAllocationMethod, error) {
	vals := map[string]IPAddressAllocationMethod{
		"dynamic":   IPAddressAllocationMethodDynamic,
		"linklayer": IPAddressAllocationMethodLinklayer,
		"other":     IPAddressAllocationMethodOther,
		"random":    IPAddressAllocationMethodRandom,
		"static":    IPAddressAllocationMethodStatic,
		"unset":     IPAddressAllocationMethodUnset,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IPAddressAllocationMethod(input)
	return &out, nil
}

type NICType string

const (
	NICTypeEOneThousand  NICType = "e1000"
	NICTypeEOneThousande NICType = "e1000e"
	NICTypePcnetThreeTwo NICType = "pcnet32"
	NICTypeVmxnet        NICType = "vmxnet"
	NICTypeVmxnetThree   NICType = "vmxnet3"
	NICTypeVmxnetTwo     NICType = "vmxnet2"
)

func PossibleValuesForNICType() []string {
	return []string{
		string(NICTypeEOneThousand),
		string(NICTypeEOneThousande),
		string(NICTypePcnetThreeTwo),
		string(NICTypeVmxnet),
		string(NICTypeVmxnetThree),
		string(NICTypeVmxnetTwo),
	}
}

func parseNICType(input string) (*NICType, error) {
	vals := map[string]NICType{
		"e1000":   NICTypeEOneThousand,
		"e1000e":  NICTypeEOneThousande,
		"pcnet32": NICTypePcnetThreeTwo,
		"vmxnet":  NICTypeVmxnet,
		"vmxnet3": NICTypeVmxnetThree,
		"vmxnet2": NICTypeVmxnetTwo,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NICType(input)
	return &out, nil
}

type OsType string

const (
	OsTypeLinux   OsType = "Linux"
	OsTypeOther   OsType = "Other"
	OsTypeWindows OsType = "Windows"
)

func PossibleValuesForOsType() []string {
	return []string{
		string(OsTypeLinux),
		string(OsTypeOther),
		string(OsTypeWindows),
	}
}

func parseOsType(input string) (*OsType, error) {
	vals := map[string]OsType{
		"linux":   OsTypeLinux,
		"other":   OsTypeOther,
		"windows": OsTypeWindows,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OsType(input)
	return &out, nil
}

type PowerOnBootOption string

const (
	PowerOnBootOptionDisabled PowerOnBootOption = "disabled"
	PowerOnBootOptionEnabled  PowerOnBootOption = "enabled"
)

func PossibleValuesForPowerOnBootOption() []string {
	return []string{
		string(PowerOnBootOptionDisabled),
		string(PowerOnBootOptionEnabled),
	}
}

func parsePowerOnBootOption(input string) (*PowerOnBootOption, error) {
	vals := map[string]PowerOnBootOption{
		"disabled": PowerOnBootOptionDisabled,
		"enabled":  PowerOnBootOptionEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PowerOnBootOption(input)
	return &out, nil
}

type SCSIControllerType string

const (
	SCSIControllerTypeBuslogic    SCSIControllerType = "buslogic"
	SCSIControllerTypeLsilogic    SCSIControllerType = "lsilogic"
	SCSIControllerTypeLsilogicsas SCSIControllerType = "lsilogicsas"
	SCSIControllerTypePvscsi      SCSIControllerType = "pvscsi"
)

func PossibleValuesForSCSIControllerType() []string {
	return []string{
		string(SCSIControllerTypeBuslogic),
		string(SCSIControllerTypeLsilogic),
		string(SCSIControllerTypeLsilogicsas),
		string(SCSIControllerTypePvscsi),
	}
}

func parseSCSIControllerType(input string) (*SCSIControllerType, error) {
	vals := map[string]SCSIControllerType{
		"buslogic":    SCSIControllerTypeBuslogic,
		"lsilogic":    SCSIControllerTypeLsilogic,
		"lsilogicsas": SCSIControllerTypeLsilogicsas,
		"pvscsi":      SCSIControllerTypePvscsi,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SCSIControllerType(input)
	return &out, nil
}

type StatusTypes string

const (
	StatusTypesConnected    StatusTypes = "Connected"
	StatusTypesDisconnected StatusTypes = "Disconnected"
	StatusTypesError        StatusTypes = "Error"
)

func PossibleValuesForStatusTypes() []string {
	return []string{
		string(StatusTypesConnected),
		string(StatusTypesDisconnected),
		string(StatusTypesError),
	}
}

func parseStatusTypes(input string) (*StatusTypes, error) {
	vals := map[string]StatusTypes{
		"connected":    StatusTypesConnected,
		"disconnected": StatusTypesDisconnected,
		"error":        StatusTypesError,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StatusTypes(input)
	return &out, nil
}

type VirtualSCSISharing string

const (
	VirtualSCSISharingNoSharing       VirtualSCSISharing = "noSharing"
	VirtualSCSISharingPhysicalSharing VirtualSCSISharing = "physicalSharing"
	VirtualSCSISharingVirtualSharing  VirtualSCSISharing = "virtualSharing"
)

func PossibleValuesForVirtualSCSISharing() []string {
	return []string{
		string(VirtualSCSISharingNoSharing),
		string(VirtualSCSISharingPhysicalSharing),
		string(VirtualSCSISharingVirtualSharing),
	}
}

func parseVirtualSCSISharing(input string) (*VirtualSCSISharing, error) {
	vals := map[string]VirtualSCSISharing{
		"nosharing":       VirtualSCSISharingNoSharing,
		"physicalsharing": VirtualSCSISharingPhysicalSharing,
		"virtualsharing":  VirtualSCSISharingVirtualSharing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualSCSISharing(input)
	return &out, nil
}
