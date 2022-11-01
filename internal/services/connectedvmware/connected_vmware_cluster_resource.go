package connectedvmware

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2020-10-01-preview/clusters"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"time"
)

type ClusterResource struct{}

var _ sdk.ResourceWithUpdate = ClusterResource{}

func (r ClusterResource) Arguments() map[string]*schema.Schema {
	return ConnectedVmwareResourceCommonSchema()
}

func (r ClusterResource) Attributes() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func (r ClusterResource) ModelObject() interface{} {
	return ClusterResource{}
}

func (r ClusterResource) ResourceType() string {
	return "azurerm_connected_vmware_cluster"
}

func (r ClusterResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			var model ConnectedVmwareResourceModel
			if err := metadata.Decode(&model); err != nil {
				return err
			}

			client := metadata.Client.ConnectedVmware.ClusterClient
			subscriptionId := metadata.Client.Account.SubscriptionId

			id := clusters.NewClusterID(subscriptionId, model.ResourceGroup, model.Name)

			existing, err := client.Get(ctx, id)
			if err != nil && !response.WasNotFound(existing.HttpResponse) {
				return fmt.Errorf("checking for presence of existing %s: %+v", id, err)
			}

			if !response.WasNotFound(existing.HttpResponse) {
				return metadata.ResourceRequiresImport(r.ResourceType(), id)
			}

			props := clusters.ClusterProperties{
				InventoryItemId: &model.InventoryItemId,
				MoRefId:         &model.MoRefId,
				VCenterId:       &model.VCenterId,
			}

			cluster := clusters.Cluster{
				ExtendedLocation: &clusters.ExtendedLocation{
					Name: &model.ExtendedLocation.Name,
					Type: &model.ExtendedLocation.Type,
				},
				Kind:     &model.Kind,
				Location: model.Location,
				Tags:     &model.Tags,
			}
			cluster.Properties = props

			if _, err := client.Create(ctx, id, cluster); err != nil {
				return fmt.Errorf("creating %s: %+v", id, err)
			}

			metadata.SetID(id)
			return nil
		},
	}
}

func (r ClusterResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.ConnectedVmware.ClusterClient

			id, err := clusters.ParseClusterID(metadata.ResourceData.Id())
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
					Name:          id.ClusterName,
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

func (r ClusterResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.ConnectedVmware.ClusterClient
			id, err := clusters.ParseClusterID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			metadata.Logger.Infof("deleting %s", *id)

			if resp, err := client.Delete(ctx, *id, clusters.DefaultDeleteOperationOptions()); err != nil {
				if !response.WasNotFound(resp.HttpResponse) {
					return fmt.Errorf("deleting %s: %+v", *id, err)
				}
			}
			return nil
		},
	}
}

func (r ClusterResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return clusters.ValidateClusterID
}

func (r ClusterResource) Update() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.ConnectedVmware.ClusterClient
			id, err := clusters.ParseClusterID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			var state ConnectedVmwareResourceModel
			if err := metadata.Decode(&state); err != nil {
				return fmt.Errorf("decoding %+v", err)
			}

			if metadata.ResourceData.HasChangesExcept("name", "resource_group_name", "location") {
				//props := clusters.Cluster{
				//	ExtendedLocation: &clusters.ExtendedLocation{
				//		Name: &state.ExtendedLocation.Name,
				//		Type: &state.ExtendedLocation.Type,
				//	},
				//	Kind: &state.Kind,
				//	Properties: clusters.ClusterProperties{
				//		InventoryItemId: &state.InventoryItemId,
				//		MoRefId:         &state.MoRefId,
				//		VCenterId:       &state.VCenterId,
				//	},
				//	Tags: &state.Tags,
				//}

				patch := clusters.ResourcePatch{
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
