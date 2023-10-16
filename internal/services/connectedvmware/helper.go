package connectedvmware

import (
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonschema"
	"github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2023-10-01/virtualmachineinstances"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/validation"
)

//type ConnectedVmwareResourceModel struct {
//	Name             string                  `tfschema:"name"`
//	ResourceGroup    string                  `tfschema:"resource_group_name"`
//	ExtendedLocation []ExtendedLocationModel `tfschema:"extended_location"`
//	Kind             string                  `tfschema:"kind"`
//	Location         string                  `tfschema:"location"`
//	InventoryItemId  string                  `tfschema:"inventory_item_id"`
//	MoRefId          string                  `tfschema:"mo_ref_id"`
//	VCenterId        string                  `tfschema:"vcenter_id"`
//	Tags             map[string]string       `tfschema:"tags"`
//}

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
						string(virtualmachineinstances.IPAddressAllocationMethodUnset),
						string(virtualmachineinstances.IPAddressAllocationMethodStatic),
						string(virtualmachineinstances.IPAddressAllocationMethodDynamic),
						string(virtualmachineinstances.IPAddressAllocationMethodLinklayer),
						string(virtualmachineinstances.IPAddressAllocationMethodOther),
						string(virtualmachineinstances.IPAddressAllocationMethodRandom),
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

func ExpandExtendedLocation(input []ExtendedLocationModel) (*string, *string) {
	if len(input) == 0 {
		return nil, nil
	}
	v := input[0]

	var locationName string
	var locationType string
	if v.Type != "" {
		locationType = v.Type
	}

	if v.Name != "" {
		locationName = v.Name
	}

	return &locationName, &locationType
}

func FlattenExtendedLocation(locationName, locationType *string) []ExtendedLocationModel {
	if locationName == nil && locationType == nil {
		return nil
	}

	return []ExtendedLocationModel{
		{
			Name: *locationName,
			Type: *locationType,
		},
	}
}

func ExpandCredential(input []CredentialModel) (*string, *string) {
	if len(input) == 0 {
		return nil, nil
	}
	v := input[0]

	var password string
	var username string
	if v.Password != "" {
		password = v.Password
	}

	if v.Username != "" {
		username = v.Username
	}

	return &username, &password
}
