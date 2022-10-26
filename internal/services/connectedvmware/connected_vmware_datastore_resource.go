package connectedvmware

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonschema"
	"github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2020-10-01-preview/datastores"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/validation"
	"time"
)

type DatastoreResource struct{}

var _ sdk.ResourceWithUpdate = DatastoreResource{}

type DatastoreResourceModel struct {
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

func (r DatastoreResource) Arguments() map[string]*schema.Schema {
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

func (r DatastoreResource) Attributes() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func (r DatastoreResource) ModelObject() interface{} {
	return DatastoreResource{}
}

func (r DatastoreResource) ResourceType() string {
	return "azurerm_connected_vmware_datastore"
}

func (r DatastoreResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			var model DatastoreResourceModel
			if err := metadata.Decode(&model); err != nil {
				return err
			}

			client := metadata.Client.ConnectedVmware.DataStoresClient
			subscriptionId := metadata.Client.Account.SubscriptionId

			id := datastores.NewDataStoreID(subscriptionId, model.ResourceGroup, model.Name)

			existing, err := client.Get(ctx, id)
			if err != nil && !response.WasNotFound(existing.HttpResponse) {
				return fmt.Errorf("checking for presence of existing %s: %+v", id, err)
			}

			if !response.WasNotFound(existing.HttpResponse) {
				return metadata.ResourceRequiresImport(r.ResourceType(), id)
			}

			props := datastores.DatastoreProperties{
				InventoryItemId: &model.InventoryItemId,
				MoRefId:         &model.MoRefId,
				VCenterId:       &model.VCenterId,
			}

			datastore := datastores.Datastore{
				ExtendedLocation: &datastores.ExtendedLocation{
					Name: &model.ExtendedLocation.Name,
					Type: &model.ExtendedLocation.Type,
				},
				Kind:     &model.Kind,
				Location: model.Location,
				Tags:     &model.Tags,
			}
			datastore.Properties = props

			if _, err := client.Create(ctx, id, datastore); err != nil {
				return fmt.Errorf("creating %s: %+v", id, err)
			}

			metadata.SetID(id)
			return nil
		},
	}
}

func (r DatastoreResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.ConnectedVmware.DataStoresClient

			id, err := datastores.ParseDataStoreID(metadata.ResourceData.Id())
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

				state := DatastoreResourceModel{
					Name:          id.DatastoreName,
					ResourceGroup: id.ResourceGroupName,
					Location:      model.Location,
				}

				if model.ExtendedLocation != nil {
					state.ExtendedLocation = ExtendedLocationModel{
						Name: *model.ExtendedLocation.Name,
						Type: *model.ExtendedLocation.Type,
					}
				}

				if model.Kind != nil {
					state.Kind = *model.Kind
				}

				if props.InventoryItemId != nil {
					state.InventoryItemId = *props.InventoryItemId
				}

				if props.MoRefId != nil {
					state.MoRefId = *props.MoRefId
				}

				if props.VCenterId != nil {
					state.VCenterId = *props.VCenterId
				}

				if model.Tags != nil {
					state.Tags = *model.Tags
				}

				return metadata.Encode(&state)
			}
			return nil
		},
	}
}

func (r DatastoreResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.ConnectedVmware.DataStoresClient
			id, err := datastores.ParseDataStoreID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			metadata.Logger.Infof("deleting %s", *id)

			if resp, err := client.Delete(ctx, *id, datastores.DefaultDeleteOperationOptions()); err != nil {
				if !response.WasNotFound(resp.HttpResponse) {
					return fmt.Errorf("deleting %s: %+v", *id, err)
				}
			}
			return nil
		},
	}
}

func (r DatastoreResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return datastores.ValidateDataStoreID
}

func (r DatastoreResource) Update() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.ConnectedVmware.DataStoresClient
			id, err := datastores.ParseDataStoreID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			var state DatastoreResourceModel
			if err := metadata.Decode(&state); err != nil {
				fmt.Errorf("decoding %+v", err)
			}

			if metadata.ResourceData.HasChangesExcept("name", "resource_group_model", "location") {
				patch := datastores.ResourcePatch{
					Tags: &state.Tags,
				}
				if _, err := client.Update(ctx, *id, patch); err != nil {
					return fmt.Errorf("updating %s: %+v", *id, err)
				}
			}
			return nil
		},
	}
}
