package connectedvmware

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2020-10-01-preview/datastores"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"time"
)

type DatastoreResource struct{}

var _ sdk.ResourceWithUpdate = DatastoreResource{}

func (r DatastoreResource) Arguments() map[string]*schema.Schema {
	return ConnectedVmwareResourceCommonSchema()
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
			var model ConnectedVmwareResourceModel
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

				state := ConnectedVmwareResourceModel{
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

			var state ConnectedVmwareResourceModel
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
