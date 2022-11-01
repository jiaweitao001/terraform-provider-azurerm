package connectedvmware

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2020-10-01-preview/resourcepools"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"time"
)

type ResourcepoolResource struct{}

var _ sdk.ResourceWithUpdate = ResourcepoolResource{}

func (r ResourcepoolResource) Arguments() map[string]*schema.Schema {
	return ConnectedVmwareResourceCommonSchema()
}

func (r ResourcepoolResource) Attributes() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func (r ResourcepoolResource) ModelObject() interface{} {
	return ResourcepoolResource{}
}

func (r ResourcepoolResource) ResourceType() string {
	return "azurerm_connected_vmware_resourcepool"
}

func (r ResourcepoolResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			var model ConnectedVmwareResourceModel
			if err := metadata.Decode(&model); err != nil {
				return err
			}

			client := metadata.Client.ConnectedVmware.ResourcepoolClient
			subscriptionId := metadata.Client.Account.SubscriptionId

			id := resourcepools.NewResourcePoolID(subscriptionId, model.ResourceGroup, model.Name)

			existing, err := client.Get(ctx, id)
			if err != nil && !response.WasNotFound(existing.HttpResponse) {
				return fmt.Errorf("checking for presence of existing %s: %+v", id, err)
			}

			if !response.WasNotFound(existing.HttpResponse) {
				return metadata.ResourceRequiresImport(r.ResourceType(), id)
			}

			props := resourcepools.ResourcePoolProperties{
				InventoryItemId: &model.InventoryItemId,
				MoRefId:         &model.MoRefId,
				VCenterId:       &model.VCenterId,
			}

			resourcepool := resourcepools.ResourcePool{
				ExtendedLocation: &resourcepools.ExtendedLocation{
					Name: &model.ExtendedLocation.Name,
					Type: &model.ExtendedLocation.Type,
				},
				Kind:     &model.Kind,
				Location: model.Location,
				Tags:     &model.Tags,
			}
			resourcepool.Properties = props

			if _, err := client.Create(ctx, id, resourcepool); err != nil {
				return fmt.Errorf("creating %s: %+v", id, err)
			}

			metadata.SetID(id)
			return nil

		},
	}
}

func (r ResourcepoolResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.ConnectedVmware.ResourcepoolClient

			id, err := resourcepools.ParseResourcePoolID(metadata.ResourceData.Id())
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
					Name:          id.ResourcePoolName,
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

func (r ResourcepoolResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.ConnectedVmware.ResourcepoolClient
			id, err := resourcepools.ParseResourcePoolID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			metadata.Logger.Infof("deleting %s", *id)

			if resp, err := client.Delete(ctx, *id, resourcepools.DefaultDeleteOperationOptions()); err != nil {
				if !response.WasNotFound(resp.HttpResponse) {
					return fmt.Errorf("deleting %s: %+v", *id, err)
				}
			}
			return nil
		},
	}
}

func (r ResourcepoolResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return resourcepools.ValidateResourcePoolID
}

func (r ResourcepoolResource) Update() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.ConnectedVmware.ResourcepoolClient
			id, err := resourcepools.ParseResourcePoolID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			var state ConnectedVmwareResourceModel
			if err := metadata.Decode(&state); err != nil {
				return fmt.Errorf("decoding %+v", err)
			}

			if metadata.ResourceData.HasChangesExcept("name", "resource_group_name", "location") {
				patch := resourcepools.ResourcePatch{
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
