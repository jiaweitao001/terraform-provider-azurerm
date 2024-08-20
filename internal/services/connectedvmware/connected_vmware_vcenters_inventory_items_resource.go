package connectedvmware

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2023-10-01/inventoryitems"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/validation"
)

type VcentersInventoryItemsResource struct{}

type VcentersInventoryItemsResourceModel struct {
	Name              string                    `tfschema:"name"`
	ResourceGroup     string                    `tfschema:"resource_group_name"`
	VcenterName       string                    `tfschema:"vcenter_name"`
	InventoryType     string                    `tfschema:"inventory_type"`
	ManagedResourceId string                    `tfschema:"managed_resource_id"`
	MoName            string                    `tfschema:"mo_name"`
	MoRefId           string                    `tfschema:"mo_ref_id"`
	Parent            InventoryItemDetailsModel `tfschema:"parent"`
	CapacityGB        int64                     `tfschema:"capacity_gb"`
	FreeSpaceGB       int64                     `tfschema:"free_space_gb"`
	FolderPath        string                    `tfschema:"folder_path"`
	MemorySizeGB      int64                     `tfschema:"memory_size_gb"`
	NumCPUs           int64                     `tfschema:"num_cpus"`
	NumCoresPerSocket int64                     `tfschema:"num_cores_per_socket"`
	OsName            string                    `tfschema:"os_name"`
	OsType            string                    `tfschema:"os_type"`
	Host              InventoryItemDetailsModel `tfschema:"host"`
	InstanceUuid      string                    `tfschema:"instance_uuid"`
	IpAddresses       []string                  `tfschema:"ip_addresses"`
	ResourcePool      InventoryItemDetailsModel `tfschema:"resource_pool"`
	SmbiosUuid        string                    `tfschema:"smbios_uuid"`
}

type InventoryItemDetailsModel struct {
	InventoryItemId string `tfschema:"inventory_item_id"`
	MoName          string `tfschema:"mo_name"`
}

func (r VcentersInventoryItemsResource) Arguments() map[string]*schema.Schema {
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

		"vcenter_name": {
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"inventory_type": {
			Type:     pluginsdk.TypeString,
			Required: true,
			ForceNew: true,
			ValidateFunc: validation.StringInSlice([]string{
				string(inventoryitems.InventoryTypeCluster),
				string(inventoryitems.InventoryTypeDatastore),
				string(inventoryitems.InventoryTypeHost),
				string(inventoryitems.InventoryTypeResourcePool),
				string(inventoryitems.InventoryTypeVirtualMachineTemplate),
				string(inventoryitems.InventoryTypeVirtualMachine),
				string(inventoryitems.InventoryTypeVirtualNetwork),
			}, false),
		},

		"managed_resource_id": {
			Type:         pluginsdk.TypeString,
			Optional:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"mo_name": {
			Type:         pluginsdk.TypeString,
			Optional:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"mo_ref_id": {
			Type:         pluginsdk.TypeString,
			Optional:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"parent": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			ForceNew: true,
			MaxItems: 1,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*schema.Schema{
					"inventory_item_id": {
						Type:         pluginsdk.TypeString,
						Required:     true,
						ValidateFunc: validation.StringIsNotEmpty,
					},

					"mo_name": {
						Type:         pluginsdk.TypeString,
						Required:     true,
						ValidateFunc: validation.StringIsNotEmpty,
					},
				},
			},
		},

		"capacity_gb": {
			Type:     pluginsdk.TypeInt,
			Optional: true,
			ForceNew: true,
		},

		"free_space_gb": {
			Type:     pluginsdk.TypeInt,
			Optional: true,
			ForceNew: true,
		},

		"folder_path": {
			Type:         pluginsdk.TypeString,
			Optional:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"memory_size_gb": {
			Type:     pluginsdk.TypeInt,
			Optional: true,
			ForceNew: true,
		},

		"num_cpus": {
			Type:     pluginsdk.TypeInt,
			Optional: true,
			ForceNew: true,
		},

		"num_cores_per_socket": {
			Type:     pluginsdk.TypeInt,
			Optional: true,
			ForceNew: true,
		},

		"os_name": {
			Type:         pluginsdk.TypeString,
			Optional:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"os_type": {
			Type:     pluginsdk.TypeString,
			Optional: true,
			ForceNew: true,
			ValidateFunc: validation.StringInSlice([]string{
				string(inventoryitems.OsTypeWindows),
				string(inventoryitems.OsTypeLinux),
				string(inventoryitems.OsTypeOther),
			}, false),
		},

		"host": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			ForceNew: true,
			MaxItems: 1,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*schema.Schema{
					"inventory_item_id": {
						Type:         pluginsdk.TypeString,
						Required:     true,
						ValidateFunc: validation.StringIsNotEmpty,
					},

					"mo_name": {
						Type:         pluginsdk.TypeString,
						Required:     true,
						ValidateFunc: validation.StringIsNotEmpty,
					},
				},
			},
		},

		"instance_uuid": {
			Type:         pluginsdk.TypeString,
			Optional:     true,
			ForceNew:     true,
			ValidateFunc: validation.IsUUID,
		},

		"ip_addresses": {
			Type:         pluginsdk.TypeString,
			Optional:     true,
			ForceNew:     true,
			ValidateFunc: validation.IsIPAddress,
		},

		"resource_pool": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			ForceNew: true,
			MaxItems: 1,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*schema.Schema{
					"inventory_item_id": {
						Type:         pluginsdk.TypeString,
						Required:     true,
						ValidateFunc: validation.StringIsNotEmpty,
					},

					"mo_name": {
						Type:         pluginsdk.TypeString,
						Required:     true,
						ValidateFunc: validation.StringIsNotEmpty,
					},
				},
			},
		},

		"smbios_uuid": {
			Type:         pluginsdk.TypeString,
			Optional:     true,
			ForceNew:     true,
			ValidateFunc: validation.IsUUID,
		},
	}
}

func (r VcentersInventoryItemsResource) Attributes() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func (r VcentersInventoryItemsResource) ModelObject() interface{} {
	return &VcentersInventoryItemsResourceModel{}
}

func (r VcentersInventoryItemsResource) ResourceType() string {
	return "azurerm_vcenters_inventory_items"
}

func (r VcentersInventoryItemsResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			var model VcentersInventoryItemsResourceModel
			if err := metadata.Decode(&model); err != nil {
				return err
			}

			client := metadata.Client.ConnectedVmware.InventoryItemsClient
			subscriptionId := metadata.Client.Account.SubscriptionId

			id := inventoryitems.NewInventoryItemID(subscriptionId, model.ResourceGroup, model.VcenterName, model.Name)

			existing, err := client.Get(ctx, id)
			if err != nil && !response.WasNotFound(existing.HttpResponse) {
				return fmt.Errorf("checking for presence of existing %s: %+v", id, err)
			}

			if !response.WasNotFound(existing.HttpResponse) {
				return metadata.ResourceRequiresImport(r.ResourceType(), id)
			}

			idStr := id.ID()
			inventoryItems := inventoryitems.InventoryItem{
				Id:   &idStr,
				Name: &model.Name,
			}
			if model.InventoryType == string(inventoryitems.InventoryTypeCluster) {
				itemProperties := inventoryitems.ClusterInventoryItem{
					ManagedResourceId: &model.ManagedResourceId,
					MoName:            &model.MoName,
					MoRefId:           &model.MoRefId,
				}
				inventoryItems.Properties = itemProperties
			}

			if model.InventoryType == string(inventoryitems.InventoryTypeResourcePool) {
				itemProperties := inventoryitems.ResourcePoolInventoryItem{
					Parent: &inventoryitems.InventoryItemDetails{
						InventoryItemId: &model.Parent.InventoryItemId,
						MoName:          &model.Parent.MoName,
					},
					ManagedResourceId: &model.ManagedResourceId,
					MoName:            &model.MoName,
					MoRefId:           &model.MoRefId,
				}
				inventoryItems.Properties = itemProperties
			}

			if model.InventoryType == string(inventoryitems.InventoryTypeHost) {
				itemProperties := inventoryitems.HostInventoryItem{
					Parent: &inventoryitems.InventoryItemDetails{
						InventoryItemId: &model.Parent.InventoryItemId,
						MoName:          &model.Parent.MoName,
					},
					ManagedResourceId: &model.ManagedResourceId,
					MoName:            &model.MoName,
					MoRefId:           &model.MoRefId,
				}
				inventoryItems.Properties = itemProperties
			}

			if model.InventoryType == string(inventoryitems.InventoryTypeDatastore) {
				itemProperties := inventoryitems.DatastoreInventoryItem{
					CapacityGB:        &model.CapacityGB,
					FreeSpaceGB:       &model.FreeSpaceGB,
					ManagedResourceId: &model.ManagedResourceId,
					MoName:            &model.MoName,
					MoRefId:           &model.MoRefId,
				}
				inventoryItems.Properties = itemProperties
			}

			if model.InventoryType == string(inventoryitems.InventoryTypeVirtualNetwork) {
				itemProperties := inventoryitems.VirtualNetworkInventoryItem{
					ManagedResourceId: &model.ManagedResourceId,
					MoName:            &model.MoName,
					MoRefId:           &model.MoRefId,
				}
				inventoryItems.Properties = itemProperties
			}

			if model.InventoryType == string(inventoryitems.InventoryTypeVirtualMachine) {
				osType := inventoryitems.OsType(model.OsType)
				itemProperties := inventoryitems.VirtualMachineInventoryItem{
					FolderPath: &model.FolderPath,
					Host: &inventoryitems.InventoryItemDetails{
						InventoryItemId: &model.Host.InventoryItemId,
						MoName:          &model.Host.MoName,
					},
					InstanceUuid: &model.InstanceUuid,
					IPAddresses:  &model.IpAddresses,
					OsName:       &model.OsName,
					OsType:       &osType,
					ResourcePool: &inventoryitems.InventoryItemDetails{
						InventoryItemId: &model.ResourcePool.InventoryItemId,
						MoName:          &model.ResourcePool.MoName,
					},
					SmbiosUuid:        &model.SmbiosUuid,
					ManagedResourceId: &model.ManagedResourceId,
					MoName:            &model.MoName,
					MoRefId:           &model.MoRefId,
				}
				inventoryItems.Properties = itemProperties
			}

			if model.InventoryType == string(inventoryitems.InventoryTypeVirtualMachineTemplate) {
				osType := inventoryitems.OsType(model.OsType)
				itemProperties := inventoryitems.VirtualMachineTemplateInventoryItem{
					FolderPath:        &model.FolderPath,
					MemorySizeMB:      &model.MemorySizeGB,
					NumCPUs:           &model.NumCPUs,
					NumCoresPerSocket: &model.NumCoresPerSocket,
					OsName:            &model.OsName,
					OsType:            &osType,
					ManagedResourceId: &model.ManagedResourceId,
					MoName:            &model.MoName,
					MoRefId:           &model.MoRefId,
				}
				inventoryItems.Properties = itemProperties
			}

			if _, err := client.Create(ctx, id, inventoryItems); err != nil {
				return fmt.Errorf("creating %s: %+v", id, err)
			}

			metadata.SetID(id)
			return nil
		},
	}
}

func (r VcentersInventoryItemsResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.ConnectedVmware.InventoryItemsClient

			id, err := inventoryitems.ParseInventoryItemID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			resp, err := client.Get(ctx, *id)
			if err != nil {
				if response.WasNotFound(resp.HttpResponse) {
					return metadata.MarkAsGone(id)
				}
				return fmt.Errorf("reading %s： %+v", *id, err)
			}

			if model := resp.Model; model != nil {
				props := model.Properties

				state := VcentersInventoryItemsResourceModel{
					Name:          id.InventoryItemName,
					ResourceGroup: id.ResourceGroupName,
					VcenterName:   id.VCenterName,
				}
				if clusterInventoryItem, ok := props.(inventoryitems.ClusterInventoryItem); ok {
					state.InventoryType = string(inventoryitems.InventoryTypeCluster)

					if clusterInventoryItem.ManagedResourceId != nil {
						state.ManagedResourceId = *clusterInventoryItem.ManagedResourceId
					}

					if clusterInventoryItem.MoName != nil {
						state.MoName = *clusterInventoryItem.MoName
					}

					if clusterInventoryItem.MoRefId != nil {
						state.MoRefId = *clusterInventoryItem.MoRefId
					}
				}

				if resourcePoolInventoryItem, ok := props.(inventoryitems.ResourcePoolInventoryItem); ok {
					state.InventoryType = string(inventoryitems.InventoryTypeResourcePool)

					if resourcePoolInventoryItem.ManagedResourceId != nil {
						state.ManagedResourceId = *resourcePoolInventoryItem.ManagedResourceId
					}

					if resourcePoolInventoryItem.MoName != nil {
						state.MoName = *resourcePoolInventoryItem.MoName
					}

					if resourcePoolInventoryItem.MoRefId != nil {
						state.MoRefId = *resourcePoolInventoryItem.MoRefId
					}

					if resourcePoolInventoryItem.Parent != nil {
						state.Parent = InventoryItemDetailsModel{
							InventoryItemId: *resourcePoolInventoryItem.Parent.InventoryItemId,
							MoName:          *resourcePoolInventoryItem.Parent.MoName,
						}
					}
				}

				if hostInventoryItem, ok := props.(inventoryitems.HostInventoryItem); ok {
					state.InventoryType = string(inventoryitems.InventoryTypeHost)

					if hostInventoryItem.ManagedResourceId != nil {
						state.ManagedResourceId = *hostInventoryItem.ManagedResourceId
					}

					if hostInventoryItem.MoName != nil {
						state.MoName = *hostInventoryItem.MoName
					}

					if hostInventoryItem.MoRefId != nil {
						state.MoRefId = *hostInventoryItem.MoRefId
					}

					if hostInventoryItem.Parent != nil {
						state.Parent = InventoryItemDetailsModel{
							InventoryItemId: *hostInventoryItem.Parent.InventoryItemId,
							MoName:          *hostInventoryItem.Parent.MoName,
						}
					}
				}

				if datastoreInventoryItem, ok := props.(inventoryitems.DatastoreInventoryItem); ok {
					state.InventoryType = string(inventoryitems.InventoryTypeDatastore)

					if datastoreInventoryItem.ManagedResourceId != nil {
						state.ManagedResourceId = *datastoreInventoryItem.ManagedResourceId
					}

					if datastoreInventoryItem.MoName != nil {
						state.MoName = *datastoreInventoryItem.MoName
					}

					if datastoreInventoryItem.MoRefId != nil {
						state.MoRefId = *datastoreInventoryItem.MoRefId
					}

					if datastoreInventoryItem.CapacityGB != nil {
						state.CapacityGB = *datastoreInventoryItem.CapacityGB
					}

					if datastoreInventoryItem.FreeSpaceGB != nil {
						state.FreeSpaceGB = *datastoreInventoryItem.FreeSpaceGB
					}
				}

				if virtualNetworkInventoryItem, ok := props.(inventoryitems.VirtualNetworkInventoryItem); ok {
					state.InventoryType = string(inventoryitems.InventoryTypeVirtualNetwork)

					if virtualNetworkInventoryItem.ManagedResourceId != nil {
						state.ManagedResourceId = *virtualNetworkInventoryItem.ManagedResourceId
					}

					if virtualNetworkInventoryItem.MoName != nil {
						state.MoName = *virtualNetworkInventoryItem.MoName
					}

					if virtualNetworkInventoryItem.MoRefId != nil {
						state.MoRefId = *virtualNetworkInventoryItem.MoRefId
					}
				}

				if virtualMachineInventoryItem, ok := props.(inventoryitems.VirtualMachineInventoryItem); ok {
					state.InventoryType = string(inventoryitems.InventoryTypeVirtualMachine)

					if virtualMachineInventoryItem.FolderPath != nil {
						state.FolderPath = *virtualMachineInventoryItem.FolderPath
					}

					if virtualMachineInventoryItem.Host != nil {
						state.Host = InventoryItemDetailsModel{
							InventoryItemId: *virtualMachineInventoryItem.Host.InventoryItemId,
							MoName:          *virtualMachineInventoryItem.Host.MoName,
						}
					}

					if virtualMachineInventoryItem.InstanceUuid != nil {
						state.InstanceUuid = *virtualMachineInventoryItem.InstanceUuid
					}

					if virtualMachineInventoryItem.IPAddresses != nil {
						state.IpAddresses = *virtualMachineInventoryItem.IPAddresses
					}

					if virtualMachineInventoryItem.OsName != nil {
						state.OsName = *virtualMachineInventoryItem.OsName
					}

					if virtualMachineInventoryItem.OsType != nil {
						state.OsType = string(*virtualMachineInventoryItem.OsType)
					}

					if virtualMachineInventoryItem.ResourcePool != nil {
						state.ResourcePool = InventoryItemDetailsModel{
							InventoryItemId: *virtualMachineInventoryItem.ResourcePool.InventoryItemId,
							MoName:          *virtualMachineInventoryItem.ResourcePool.MoName,
						}
					}

					if virtualMachineInventoryItem.SmbiosUuid != nil {
						state.SmbiosUuid = *virtualMachineInventoryItem.SmbiosUuid
					}

					if virtualMachineInventoryItem.ManagedResourceId != nil {
						state.ManagedResourceId = *virtualMachineInventoryItem.ManagedResourceId
					}

					if virtualMachineInventoryItem.MoName != nil {
						state.MoName = *virtualMachineInventoryItem.MoName
					}

					if virtualMachineInventoryItem.MoRefId != nil {
						state.MoRefId = *virtualMachineInventoryItem.MoRefId
					}
				}

				if virtualMachineTemplateInventoryItem, ok := props.(inventoryitems.VirtualMachineTemplateInventoryItem); ok {
					state.InventoryType = string(inventoryitems.InventoryTypeVirtualMachineTemplate)

					if virtualMachineTemplateInventoryItem.FolderPath != nil {
						state.FolderPath = *virtualMachineTemplateInventoryItem.FolderPath
					}

					if virtualMachineTemplateInventoryItem.MemorySizeMB != nil {
						state.MemorySizeGB = *virtualMachineTemplateInventoryItem.MemorySizeMB
					}

					if virtualMachineTemplateInventoryItem.NumCPUs != nil {
						state.NumCPUs = *virtualMachineTemplateInventoryItem.NumCPUs
					}

					if virtualMachineTemplateInventoryItem.NumCoresPerSocket != nil {
						state.NumCoresPerSocket = *virtualMachineTemplateInventoryItem.NumCoresPerSocket
					}

					if virtualMachineTemplateInventoryItem.OsName != nil {
						state.OsName = *virtualMachineTemplateInventoryItem.OsName
					}

					if virtualMachineTemplateInventoryItem.OsType != nil {
						state.OsType = string(*virtualMachineTemplateInventoryItem.OsType)
					}

					if virtualMachineTemplateInventoryItem.ManagedResourceId != nil {
						state.ManagedResourceId = *virtualMachineTemplateInventoryItem.ManagedResourceId
					}

					if virtualMachineTemplateInventoryItem.MoName != nil {
						state.MoName = *virtualMachineTemplateInventoryItem.MoName
					}

					if virtualMachineTemplateInventoryItem.MoRefId != nil {
						state.MoRefId = *virtualMachineTemplateInventoryItem.MoRefId
					}
				}

				return metadata.Encode(&state)
			}
			return nil
		},
	}
}

func (r VcentersInventoryItemsResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.ConnectedVmware.InventoryItemsClient
			id, err := inventoryitems.ParseInventoryItemID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			metadata.Logger.Infof("deleting %s", *id)

			if resp, err := client.Delete(ctx, *id); err != nil {
				if !response.WasNotFound(resp.HttpResponse) {
					return fmt.Errorf("deleting %s: %+v", *id, err)
				}
			}

			return nil
		},
	}
}

func (r VcentersInventoryItemsResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return inventoryitems.ValidateInventoryItemID
}
