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
	"github.com/hashicorp/terraform-provider-azurerm/utils"
	"time"
)

type VirtualMachineResource struct{}

var _ sdk.ResourceWithUpdate = VirtualMachineResource{}

type VirtualMachineResourceModel struct {
	ConnectedVmwareResourceProperties []ConnectedVmwareResourceModel `tfschema:"connected_vmware"`
	FirmwareType                      string                         `tfschema:"firmware_type"`
	MemorySizeMB                      int64                          `tfschema:"memory_size_MB"`
	NumCPUs                           int64                          `tfschema:"num_CPUs"`
	NumCoresPerSocket                 int64                          `tfschema:"num_cores_per_socket"`
	NetworkInterface                  []NetworkInterfaceModel        `tfschema:"network_interface"`
	AdminPassword                     string                         `tfschema:"admin_password"`
	AdminUsername                     string                         `tfschema:"admin_username"`
	ComputerName                      string                         `tfschema:"computer_name"`
	GuestId                           string                         `tfschema:"guest_id"`
	OsType                            string                         `tfschema:"os_type"`
	OsConfiguration                   []OsConfigurationModel         `tfschema:"os_configuration"`
	ClusterId                         string                         `tfschema:"cluster_id"`
	DatastoreId                       string                         `tfschema:"datastore_id"`
	HostId                            string                         `tfschema:"host_id"`
	PlacementResourcePoolId           string                         `tfschema:"placement_resource_pool_id"`
	ResourcePoolId                    string                         `tfschema:"resource_pool_id"`
	SecureBootEnabled                 bool                           `tfschema:"secure_boot_enabled"`
	SmbiosUuid                        string                         `tfschema:"smbios_uuid"`
	Disks                             []DisksModel                   `tfschema:"disks"`
	TemplateId                        string                         `tfschema:"template_id"`
}

type NetworkInterfaceModel struct {
	DeviceKey   int64            `tfschema:"device_key"`
	IpSettings  []IpSettingModel `tfschema:"ip_settings"`
	NetworkName string           `tfschema:"network_name"`
	NetworkId   string           `tfschema:"network_id"`
	NicType     string           `tfschema:"nic_type"`
	PowerOnBoot string           `tfschema:"power_on_boot"`
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

type DisksModel struct {
	ControllerKey int64  `tfschema:"controller_key"`
	DeviceKey     int64  `tfschema:"device_key"`
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

		"network_interface": networkInterfaceSchema(),

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

		"resource_pool_id": {
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
						Type:     pluginsdk.TypeInt,
						Optional: true,
					},

					"device_key": {
						Type:     pluginsdk.TypeInt,
						Optional: true,
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

		"template_id": {
			Type:         pluginsdk.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringIsNotEmpty,
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
			connectedVmwareCommonProps := model.ConnectedVmwareResourceProperties[0]

			id := virtualmachines.NewVirtualMachineID(subscriptionId, connectedVmwareCommonProps.ResourceGroup, connectedVmwareCommonProps.Name)

			existing, err := client.Get(ctx, id)
			if err != nil && !response.WasNotFound(existing.HttpResponse) {
				return fmt.Errorf("checking for presence of existing %s: %+v", id, err)
			}

			if !response.WasNotFound(existing.HttpResponse) {
				return metadata.ResourceRequiresImport(r.ResourceType(), id)
			}

			osType := virtualmachines.OsType(model.OsType)
			firmwareType := virtualmachines.FirmwareType(model.FirmwareType)
			props := virtualmachines.VirtualMachineProperties{
				CustomResourceName: utils.String(connectedVmwareCommonProps.Name),
				HardwareProfile: &virtualmachines.HardwareProfile{
					MemorySizeMB:      utils.Int64(model.MemorySizeMB),
					NumCPUs:           utils.Int64(model.NumCPUs),
					NumCoresPerSocket: utils.Int64(model.NumCoresPerSocket),
				},
				InventoryItemId: utils.String(connectedVmwareCommonProps.InventoryItemId),
				FirmwareType:    &firmwareType,
				MoRefId:         utils.String(connectedVmwareCommonProps.MoRefId),
				NetworkProfile: &virtualmachines.NetworkProfile{
					NetworkInterfaces: ExpandNetworkInterface(model.NetworkInterface),
				},
				OsProfile: &virtualmachines.OsProfile{
					AdminPassword: utils.String(model.AdminPassword),
					AdminUsername: utils.String(model.AdminUsername),
					ComputerName:  utils.String(model.ComputerName),
					OsType:        &osType,
				},
				PlacementProfile: &virtualmachines.PlacementProfile{
					ClusterId:      utils.String(model.ClusterId),
					DatastoreId:    utils.String(model.DatastoreId),
					HostId:         utils.String(model.HostId),
					ResourcePoolId: utils.String(model.PlacementResourcePoolId),
				},
				ResourcePoolId: &model.ResourcePoolId,
				SmbiosUuid:     utils.String(model.SmbiosUuid),
				StorageProfile: &virtualmachines.StorageProfile{
					Disks: ExpandDisks(model.Disks),
				},
				TemplateId: utils.String(model.TemplateId),
				VCenterId:  utils.String(connectedVmwareCommonProps.VCenterId),
			}

			if osType == virtualmachines.OsTypeWindows {
			}

			virtualMachine := virtualmachines.VirtualMachine{
				ExtendedLocation: &virtualmachines.ExtendedLocation{
					Name: utils.String(connectedVmwareCommonProps.ExtendedLocation.Name),
					Type: utils.String(connectedVmwareCommonProps.ExtendedLocation.Type),
				},
				Id:         utils.String(id.ID()),
				Kind:       utils.String(connectedVmwareCommonProps.Kind),
				Location:   connectedVmwareCommonProps.Location,
				Properties: props,
				Tags:       &connectedVmwareCommonProps.Tags,
			}

			if err := client.CreateThenPoll(ctx, id, virtualMachine); err != nil {
				return fmt.Errorf("creating %s： %+v", id, err)
			}

			metadata.SetID(id)
			return nil
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
				if response.WasNotFound(resp.HttpResponse) {
					return metadata.MarkAsGone(id)
				}
				return fmt.Errorf("reading %s: %+v", *id, err)
			}

			if model := resp.Model; model != nil {
				props := model.Properties

				state := VirtualMachineResourceModel{
					ConnectedVmwareResourceProperties: FlattenConnectedVmwareResourceProperties(model, id),
					FirmwareType:                      string(*props.FirmwareType),
					NetworkInterface:                  FlattenNetworkProfile(props.NetworkProfile),
					ResourcePoolId:                    *props.ResourcePoolId,
					SmbiosUuid:                        *props.SmbiosUuid,
					Disks:                             FlattenStorageProfile(props.StorageProfile),
					TemplateId:                        *props.TemplateId,
				}

				if props.HardwareProfile != nil {
					state.MemorySizeMB = *props.HardwareProfile.MemorySizeMB
					state.NumCPUs = *props.HardwareProfile.NumCPUs
					state.NumCoresPerSocket = *props.HardwareProfile.NumCoresPerSocket
				}

				if props.OsProfile != nil {
					state.AdminPassword = *props.OsProfile.AdminPassword
					state.AdminUsername = *props.OsProfile.AdminUsername
					state.ComputerName = *props.OsProfile.ComputerName
					state.OsType = string(*props.OsProfile.OsType)
				}

				if props.PlacementProfile != nil {
					state.ClusterId = *props.PlacementProfile.ClusterId
					state.DatastoreId = *props.PlacementProfile.DatastoreId
					state.HostId = *props.PlacementProfile.HostId
					state.ResourcePoolId = *props.PlacementProfile.ResourcePoolId
				}

				return metadata.Encode(&state)
			}
			return nil
		},
	}
}

func (r VirtualMachineResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.ConnectedVmware.VirtualMachineClient
			id, err := virtualmachines.ParseVirtualMachineID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			metadata.Logger.Infof("deleting %s", *id)

			if resp, err := client.Delete(ctx, *id, virtualmachines.DefaultDeleteOperationOptions()); err != nil {
				if !response.WasNotFound(resp.HttpResponse) {
					return fmt.Errorf("deleting %s: %+v", *id, err)
				}
			}
			return nil
		},
	}
}

func (r VirtualMachineResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return virtualmachines.ValidateVirtualMachineID
}

func (r VirtualMachineResource) Update() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.ConnectedVmware.VirtualMachineClient
			id, err := virtualmachines.ParseVirtualMachineID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			var state VirtualMachineResourceModel
			if err := metadata.Decode(&state); err != nil {
				return fmt.Errorf("decoding %+v", err)
			}

			if metadata.ResourceData.HasChangeExcept()

			client.Update()
		},
	}
}
