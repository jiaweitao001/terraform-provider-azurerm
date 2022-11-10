package connectedvmware

import (
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonschema"
	"github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2020-10-01-preview/virtualmachines"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/validation"
	"github.com/hashicorp/terraform-provider-azurerm/utils"
)

type ConnectedVmwareResourceModel struct {
	Name             string                `tfschema:"name"`
	ResourceGroup    string                `tfschema:"resource_group_name"`
	ExtendedLocation ExtendedLocationModel `tfschema:"extended_location"`
	Kind             string                `tfschema:"kind"`
	Location         string                `tfschema:"location"`
	InventoryItemId  string                `tfschema:"inventory_item_id"`
	MoRefId          string                `tfschema:"mo_ref_id"`
	VCenterId        string                `tfschema:"vcenter_id"`
	Tags             map[string]string     `tfschema:"tags"`
}

type ExtendedLocationModel struct {
	Name string `tfschema:"name"`
	Type string `tfschema:"type"`
}

func ConnectedVmwareResourceCommonSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"resource_group_name": {
			Type:     pluginsdk.TypeString,
			Required: true,
			ForceNew: true,
		},

		"location": {
			Type:     pluginsdk.TypeString,
			Required: true,
			ForceNew: true,
		},

		"extended_location": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Type:         pluginsdk.TypeString,
						Required:     true,
						ValidateFunc: validation.StringIsNotEmpty,
					},

					"type": {
						Type:         pluginsdk.TypeString,
						Required:     true,
						ValidateFunc: validation.StringIsNotEmpty,
					},
				},
			},
		},

		"kind": {
			Type:     pluginsdk.TypeString,
			Optional: true,
		},

		"inventory_item_id": {
			Type:         pluginsdk.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"mo_ref_id": {
			Type:         pluginsdk.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"vcenter_id": {
			Type:         pluginsdk.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"tags": commonschema.Tags(),
	}
}

func networkInterfaceSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Optional: true,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*schema.Schema{
				"device_key": {
					Type:     pluginsdk.TypeInt,
					Optional: true,
				},

				"ip_setting": ipSettingSchema(),

				"network_name": {
					Type:         pluginsdk.TypeString,
					Optional:     true,
					ValidateFunc: validation.StringIsNotEmpty,
				},

				"network_id": {
					Type:         pluginsdk.TypeString,
					Optional:     true,
					ValidateFunc: validation.StringIsNotEmpty,
				},

				"nic_type": {
					Type:         pluginsdk.TypeString,
					Optional:     true,
					ValidateFunc: validation.StringIsNotEmpty,
				},

				"power_on_boot": {
					Type:         pluginsdk.TypeString,
					Optional:     true,
					ValidateFunc: validation.StringIsNotEmpty,
				},
			},
		},
	}
}

func ipSettingSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Optional: true,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*schema.Schema{
				"allocation_method": {
					Type:     pluginsdk.TypeString,
					Optional: true,
					ValidateFunc: validation.StringInSlice([]string{
						string(virtualmachines.IPAddressAllocationMethodUnset),
						string(virtualmachines.IPAddressAllocationMethodStatic),
						string(virtualmachines.IPAddressAllocationMethodDynamic),
						string(virtualmachines.IPAddressAllocationMethodLinklayer),
						string(virtualmachines.IPAddressAllocationMethodOther),
						string(virtualmachines.IPAddressAllocationMethodRandom),
					}, false),
				},

				"dns_servers": {
					Type:     pluginsdk.TypeList,
					Optional: true,
					Elem: &pluginsdk.Schema{
						Type: pluginsdk.TypeString,
					},
				},

				"gateway": {
					Type:     pluginsdk.TypeList,
					Optional: true,
					Elem: &pluginsdk.Schema{
						Type: pluginsdk.TypeString,
					},
				},

				"ip_address": {
					Type:         pluginsdk.TypeString,
					Optional:     true,
					ValidateFunc: validation.StringIsNotEmpty,
				},

				"subet_mask": {
					Type:         pluginsdk.TypeString,
					Optional:     true,
					ValidateFunc: validation.StringIsNotEmpty,
				},
			},
		},
	}
}

func ExpandNetworkInterface(input []NetworkInterfaceModel) *[]virtualmachines.NetworkInterface {
	if len(input) == 0 {
		return nil
	}

	var networkInterfaces []virtualmachines.NetworkInterface

	for _, v := range input {
		var networkInterface virtualmachines.NetworkInterface
		networkInterface.DeviceKey = utils.Int64(v.DeviceKey)

		if v.NetworkId != "" {
			networkInterface.NetworkId = utils.String(v.NetworkId)
		}

		if v.NetworkName != "" {
			networkInterface.Name = utils.String(v.NetworkName)
		}

		if v.NicType != "" {
			nicType := virtualmachines.NICType(v.NicType)
			networkInterface.NicType = &nicType
		}

		if v.PowerOnBoot != "" {
			powerOnBoot := virtualmachines.PowerOnBootOption(v.PowerOnBoot)
			networkInterface.PowerOnBoot = &powerOnBoot
		}

		networkInterface.IpSettings = ExpandIpSettings(v.IpSettings)

		networkInterfaces = append(networkInterfaces, networkInterface)
	}

	return &networkInterfaces
}

func ExpandIpSettings(input []IpSettingModel) *virtualmachines.NicIPSettings {
	if len(input) == 0 {
		return nil
	}

	var settings virtualmachines.NicIPSettings

	v := input[0]
	if v.AllocationMethod != "" {
		allocationMethod := virtualmachines.IPAddressAllocationMethod(v.AllocationMethod)
		settings.AllocationMethod = &allocationMethod
	}
	if len(v.Gateway) > 0 {
		settings.Gateway = utils.StringSlice(v.Gateway)
	}
	if len(v.DnsServers) > 0 {
		settings.DnsServers = utils.StringSlice(v.DnsServers)
	}
	if v.IpAddress != "" {
		settings.IpAddress = utils.String(v.IpAddress)
	}
	if v.SubnetMask != "" {
		settings.SubnetMask = utils.String(v.SubnetMask)
	}

	return &settings
}

func ExpandDisks(input []DiskseModel) *[]virtualmachines.VirtualDisk {
	if len(input) == 0 {
		return nil
	}

	var disks []virtualmachines.VirtualDisk

	for _, v := range input {
		var disk virtualmachines.VirtualDisk

		disk.ControllerKey = utils.Int64(v.ControllerKey)

		disk.DeviceKey = utils.Int64(v.DeviceKey)

		if v.DeviceName != "" {
			disk.DeviceName = utils.String(v.DeviceName)
		}

		if v.DiskMode != "" {
			mode := virtualmachines.DiskMode(v.DiskMode)
			disk.DiskMode = &mode
		}

		disk.DiskSizeGB = utils.Int64(v.DiskSizeGB)

		if v.DiskType != "" {
			diskType := virtualmachines.DiskType(v.DiskType)
			disk.DiskType = &diskType
		}

		if v.DiskName != "" {
			disk.Name = utils.String(v.DiskName)
		}

		disk.UnitNumber = utils.Int64(v.UnitNumber)

		disks = append(disks, disk)
	}

	return &disks
}

func FlattenConnectedVmwareResourceProperties(input *virtualmachines.VirtualMachine, id *virtualmachines.VirtualMachineId) []ConnectedVmwareResourceModel {
	var props ConnectedVmwareResourceModel

	virtualMachineProps := input.Properties
	if virtualMachineProps.VCenterId != nil {
		props.VCenterId = *virtualMachineProps.VCenterId
	}

	if virtualMachineProps.MoRefId != nil {
		props.MoRefId = *virtualMachineProps.MoRefId
	}

	if virtualMachineProps.InventoryItemId != nil {
		props.InventoryItemId = *virtualMachineProps.InventoryItemId
	}

	if input.Kind != nil {
		props.Kind = *input.Kind
	}

	if input.Tags != nil {
		props.Tags = *input.Tags
	}

	if input.ExtendedLocation != nil {
		props.ExtendedLocation = ExtendedLocationModel{
			Name: *input.ExtendedLocation.Name,
			Type: *input.ExtendedLocation.Type,
		}
	}

	props.Location = input.Location

	props.Name = *input.Name

	props.ResourceGroup = id.ResourceGroupName

	return []ConnectedVmwareResourceModel{props}
}

func FlattenNetworkProfile(input *virtualmachines.NetworkProfile) []NetworkInterfaceModel {
	networkInterfaces := input.NetworkInterfaces
	if networkInterfaces == nil || len(*networkInterfaces) == 0 {
		return []NetworkInterfaceModel{}
	}

	var res []NetworkInterfaceModel
	values := *networkInterfaces

	for _, v := range values {
		var networkInterface NetworkInterfaceModel

		if v.DeviceKey != nil {
			networkInterface.DeviceKey = *v.DeviceKey
		}

		if v.IpSettings != nil {
			networkInterface.IpSettings = FlattenIpSettings(v.IpSettings)
		}

		if v.Name != nil {
			networkInterface.NetworkName = *v.Name
		}

		if v.NetworkId != nil {
			networkInterface.NetworkId = *v.NetworkId
		}

		if v.NicType != nil {
			networkInterface.NicType = string(*v.NicType)
		}

		if v.PowerOnBoot != nil {
			networkInterface.PowerOnBoot = string(*v.PowerOnBoot)
		}

		res = append(res, networkInterface)
	}

	return res
}

func FlattenIpSettings(input *virtualmachines.NicIPSettings) []IpSettingModel {
	var ipSetting IpSettingModel

	if input.Gateway != nil {
		ipSetting.Gateway = *input.Gateway
	}

	if input.IpAddress != nil {
		ipSetting.IpAddress = *input.IpAddress
	}

	if input.SubnetMask != nil {
		ipSetting.SubnetMask = *input.SubnetMask
	}

	if input.DnsServers != nil {
		ipSetting.DnsServers = *input.DnsServers
	}

	if input.AllocationMethod != nil {
		ipSetting.AllocationMethod = string(*input.AllocationMethod)
	}

	return []IpSettingModel{ipSetting}
}

func FlattenStorageProfile(input *virtualmachines.StorageProfile) []DisksModel {
	disks := input.Disks
	if input.Disks == nil {
		return []DisksModel{}
	}

	var res []DisksModel
	values := *disks

	for _, v := range values {
		var disk DisksModel

		if v.ControllerKey != nil {
			disk.ControllerKey = *v.ControllerKey
		}

		if v.DeviceKey != nil {
			disk.DeviceKey = *v.DeviceKey
		}

		if v.DeviceName != nil {
			disk.DeviceName = *v.DeviceName
		}

		if v.DiskMode != nil {
			disk.DiskMode = string(*v.DiskMode)
		}

		if v.DiskSizeGB != nil {
			disk.DiskSizeGB = *v.DiskSizeGB
		}

		if v.UnitNumber != nil {
			disk.UnitNumber = *v.UnitNumber
		}

		if v.DiskType != nil {
			disk.DiskType = string(*v.DiskType)
		}

		if v.Name != nil {
			disk.DiskName = *v.Name
		}

		res = append(res, disk)
	}

	return res
}
