package connectedvmware

import (
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonschema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/validation"
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
