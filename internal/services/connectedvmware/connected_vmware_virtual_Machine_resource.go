package connectedvmware

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2020-10-01-preview/virtualmachines"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/validation"
	"time"
)

type VirtualMachineResource struct{}

var _ sdk.ResourceWithUpdate = VirtualMachineResource{}

type VirtualMachineResourceModel struct {
	ConnectedVmwareResourceProperties ConnectedVmwareResourceModel `tfschema:"connected_vmware"`
	FirmwareType                      string                       `tfschema:"firmware_type"`
	MemorySizeMB                      int64                        `tfschema:"memory_size_MB"`
	NumCPUs                           int64                        `tfschema:"num_CPUs"`
	NumCoresPerSocket                 int64                        `tfschema:"num_cores_per_socket"`
	DeviceKey                         int64                        `tfschema:"device_key"`
	IpSettings                        IpSettingModel               `tfschema:"ip_settings"`
	NetworkName                       string                       `tfschema:"network_name"`
	NetworkId                         string                       `tfschema:"network_id"`
	NicType                           string                       `tfschema:"nic_type"`
	PowerOnBoot                       string                       `tfschema:"power_on_boot"`
	AdminPassword                     string                       `tfschema:"admin_password"`
	AdminUsername                     string                       `tfschema:"admin_username"`
	ComputerName                      string                       `tfschema:"computer_name"`
	GuestId                           string                       `tfschema:"guest_id"`
	OsType                            string                       `tfschema:"os_type"`
	OsConfiguration                   OsConfigurationModel         `tfschema:"os_configuration"`
	ClusterId                         string                       `tfschema:"cluster_id"`
	DatastoreId                       string                       `tfschema:"datastore_id"`
	HostId                            string                       `tfschema:"host_id"`
	ResourcepoolId                    string                       `tfschema:"resourcepool_id"`
	SecureBootEnabled                 bool                         `tfschema:"secure_boot_enabled"`
	SmbiosUuid                        string                       `tfschema:"smbios_uuid"`
	StorageProfile                    StorageProfileModel          `tfschema:"storage_profile"`
}

type IpSettingModel struct {
	AllocationMethod string   `tfschema:"allocation_method"`
	DnsServers       []string `tfschema:"dns_servers"`
	Gateway          []string `tfschema:"gateway"`
	IpAddress        string   `tfschema:"ip_address"`
	SubnetMask       string   `tfschema:"subnet_mask"`
}

type OsConfigurationModel struct {
	AssessmentMode string `tfschema:"assessment_mode"`
	PatchMode      string `tfschema:"patch_mode"`
}

type StorageProfileModel struct {
	ControllerKey string `tfschema:"controller_key"`
	DeviceKey     string `tfschema:"device_key"`
	DeviceName    string `tfschma:"device_name"`
	DiskMode      string `tfschema:"disk_mode"`
	DiskSizeGB    int64  `tfschema:"disk_size_GB"`
	DiskType      string `tfschema:"disk_type"`
	DiskName      string `tfschema:"disk_name"`
	UnitNumber    int64  `tfschema:"unit_number"`
}

func (r VirtualMachineResource) Arguments() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"connected_vmware": {
			Type:     pluginsdk.TypeList,
			Required: true,
			MaxItems: 1,
			MinItems: 1,
			Elem: &pluginsdk.Resource{
				Schema: ConnectedVmwareResourceCommonSchema(),
			},
		},

		"firmware_type": {
			Type:     pluginsdk.TypeString,
			Optional: true,
			ValidateFunc: validation.StringInSlice([]string{
				string(virtualmachines.FirmwareTypeBios),
				string(virtualmachines.FirmwareTypeEfi),
			}, false),
		},

		"memory_size_MB": {
			Type:     pluginsdk.TypeInt,
			Optional: true,
		},

		"num_CPUs": {
			Type:     pluginsdk.TypeInt,
			Optional: true,
		},

		"num_cores_per_socket": {
			Type:     pluginsdk.TypeInt,
			Optional: true,
		},

		"device_key": {
			Type:     pluginsdk.TypeInt,
			Optional: true,
		},

		"ip_settings": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*schema.Schema{
					"allocation_method": {
						Type:     pluginsdk.TypeString,
						Optional: true,
						ValidateFunc: validation.StringInSlice([]string{
							string(virtualmachines.IPAddressAllocationMethodDynamic),
							string(virtualmachines.IPAddressAllocationMethodLinklayer),
							string(virtualmachines.IPAddressAllocationMethodOther),
							string(virtualmachines.IPAddressAllocationMethodRandom),
							string(virtualmachines.IPAddressAllocationMethodStatic),
							string(virtualmachines.IPAddressAllocationMethodUnset),
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

					"subnet_mask": {
						Type:         pluginsdk.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringIsNotEmpty,
					},
				},
			},
		},

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
			Type:     pluginsdk.TypeString,
			Optional: true,
			ValidateFunc: validation.StringInSlice([]string{
				string(virtualmachines.PowerOnBootOptionEnabled),
				string(virtualmachines.PowerOnBootOptionDisabled),
			}, false),
		},

		"admin_password": {
			Type:         pluginsdk.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"admin_username": {
			Type:         pluginsdk.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"computer_name": {
			Type:         pluginsdk.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"guest_id": {
			Type:         pluginsdk.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"os_type": {
			Type:     pluginsdk.TypeString,
			Optional: true,
			ValidateFunc: validation.StringInSlice([]string{
				string(virtualmachines.OsTypeLinux),
				string(virtualmachines.OsTypeWindows),
			}, false),
		},

		"os_configuration": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*schema.Schema{
					"assessment_mode": {
						Type:         pluginsdk.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringIsNotEmpty,
					},

					"patch_mode": {
						Type:         pluginsdk.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringIsNotEmpty,
					},
				},
			},
		},

		"cluster_id": {
			Type:         pluginsdk.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"datastore_id": {
			Type:         pluginsdk.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"host_id": {
			Type:         pluginsdk.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"resourcepool_id": {
			Type:         pluginsdk.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"secure_boot_enabled": {
			Type:     pluginsdk.TypeBool,
			Optional: true,
		},

		"smbios_uuid": {
			Type:         pluginsdk.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"storage_profile": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*schema.Schema{
					"controller_key": {
						Type:         pluginsdk.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringIsNotEmpty,
					},

					"device_key": {
						Type:         pluginsdk.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringIsNotEmpty,
					},

					"device_name": {
						Type:         pluginsdk.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringIsNotEmpty,
					},

					"disk_mode": {
						Type:     pluginsdk.TypeString,
						Optional: true,
						ValidateFunc: validation.StringInSlice([]string{
							string(virtualmachines.DiskModeIndependentNonpersistent),
							string(virtualmachines.DiskModeIndependentPersistent),
							string(virtualmachines.DiskModePersistent),
						}, false),
					},

					"disk_size_GB": {
						Type:     pluginsdk.TypeInt,
						Optional: true,
					},

					"disk_type": {
						Type:     pluginsdk.TypeString,
						Optional: true,
						ValidateFunc: validation.StringInSlice([]string{
							string(virtualmachines.DiskTypeFlat),
							string(virtualmachines.DiskTypePmem),
							string(virtualmachines.DiskTypeRawphysical),
							string(virtualmachines.DiskTypeRawvirtual),
							string(virtualmachines.DiskTypeSesparse),
							string(virtualmachines.DiskTypeSparse),
							string(virtualmachines.DiskTypeUnknown),
						}, false),
					},

					"disk_name": {
						Type:         pluginsdk.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringIsNotEmpty,
					},

					"unit_number": {
						Type:     pluginsdk.TypeInt,
						Optional: true,
					},
				},
			},
		},
	}
}

func (r VirtualMachineResource) Attributes() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func (r VirtualMachineResource) ModelObject() interface{} {
	return VirtualMachineResource{}
}

func (r VirtualMachineResource) ResourceType() string {
	return "azurerm_connected_vmware_virtual_machine"
}

func (r VirtualMachineResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			var model VirtualMachineResourceModel
			if err := metadata.Decode(&model); err != nil {
				return err
			}

			client := metadata.Client.ConnectedVmware.VirtualMachineClient
			subscriptionId := metadata.Client.Account.SubscriptionId

			id := virtualmachines.NewVirtualMachineID(subscriptionId, model.ConnectedVmwareResourceProperties.ResourceGroup, model.ConnectedVmwareResourceProperties.Name)

			existing, err := client.Get(ctx, id)
			if err != nil && !response.WasNotFound(existing.HttpResponse) {
				return fmt.Errorf("checking for presence of existing %s: %+v", id, err)
			}

			if !response.WasNotFound(existing.HttpResponse) {
				return metadata.ResourceRequiresImport(r.ResourceType(), id)
			}

		},
	}
}

func (r VirtualMachineResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.ConnectedVmware.VirtualMachineClient

			id, err := virtualmachines.ParseVirtualMachineID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			resp, err := client.Get(ctx, *id)
			if err != nil {

			}
		},
	}
}

func (r VirtualMachineResource) Delete() sdk.ResourceFunc {
	//TODO implement me
	panic("implement me")
}

func (r VirtualMachineResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	//TODO implement me
	panic("implement me")
}

func (r VirtualMachineResource) Update() sdk.ResourceFunc {
	//TODO implement me
	panic("implement me")
}
