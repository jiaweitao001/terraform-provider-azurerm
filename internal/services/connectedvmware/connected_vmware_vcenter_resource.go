package connectedvmware

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonschema"
	"github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2020-10-01-preview/vcenters"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/validation"
	"time"
)

type VcenterResource struct{}

var _ sdk.ResourceWithUpdate = VcenterResource{}

type VcenterResourceModel struct {
	Name             string                  `tfschema:"name"`
	ResourceGroup    string                  `tfschema:"resource_group_name"`
	ExtendedLocation []ExtendedLocationModel `tfschema:"extended_location"`
	Kind             string                  `tfschema:"kind"`
	Location         string                  `tfschema:"location"`
	Credential       []CredentialModel       `tfschema:"credential"`
	Fqdn             string                  `tfschema:"fqdn"`
	Port             int64                   `tfschema:"port"`
	Tags             map[string]string       `tfschema:"tags"`
}

type CredentialModel struct {
	Password string `tfschema:"password"`
	Username string `tfschema:"username"`
}

func (r VcenterResource) Arguments() map[string]*schema.Schema {
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

		"fqdn": {
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"extended_location": {
			Type:     pluginsdk.TypeList,
			Required: true,
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

		"credential": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*schema.Schema{
					"username": {
						Type:         pluginsdk.TypeString,
						Required:     true,
						ValidateFunc: validation.StringIsNotEmpty,
					},

					"password": {
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

		"port": {
			Type:         pluginsdk.TypeString,
			Optional:     true,
			ValidateFunc: validation.IsPortNumber,
		},

		"tags": commonschema.Tags(),
	}
}

func (r VcenterResource) Attributes() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func (r VcenterResource) ModelObject() interface{} {
	return &VcenterResourceModel{}
}

func (r VcenterResource) ResourceType() string {
	return "azurerm_connected_vmware_vcenter"
}

func (r VcenterResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			var model VcenterResourceModel
			if err := metadata.Decode(&model); err != nil {
				return err
			}

			client := metadata.Client.ConnectedVmware.VcenterClient
			subscriptionId := metadata.Client.Account.SubscriptionId

			id := vcenters.NewVCenterID(subscriptionId, model.ResourceGroup, model.Name)

			existing, err := client.Get(ctx, id)
			if err != nil && !response.WasNotFound(existing.HttpResponse) {
				return fmt.Errorf("checking for presence of existing %s: %+v", id, err)
			}

			if !response.WasNotFound(existing.HttpResponse) {
				return metadata.ResourceRequiresImport(r.ResourceType(), id)
			}

			props := vcenters.VCenterProperties{
				Fqdn: model.Fqdn,
			}

			username, password := ExpandCredential(model.Credential)
			props.Credentials = &vcenters.VICredential{
				Password: password,
				Username: username,
			}

			if model.Port != 0 {
				props.Port = &model.Port
			}

			vcenter := vcenters.VCenter{
				Location: model.Location,
			}

			locationName, locationType := ExpandExtendedLocation(model.ExtendedLocation)
			vcenter.ExtendedLocation = &vcenters.ExtendedLocation{
				Name: locationName,
				Type: locationType,
			}

			if model.Kind != "" {
				vcenter.Kind = &model.Kind
			}

			if model.Tags != nil {
				vcenter.Tags = &model.Tags
			}

			vcenter.Properties = props

			if _, err := client.Create(ctx, id, vcenter); err != nil {
				return fmt.Errorf("creating %s: %+v", id, err)
			}

			metadata.SetID(id)
			return nil
		},
	}
}

func (r VcenterResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.ConnectedVmware.VcenterClient

			id, err := vcenters.ParseVCenterID(metadata.ResourceData.Id())
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

				state := VcenterResourceModel{
					Name:          id.VcenterName,
					ResourceGroup: id.ResourceGroupName,
					Location:      model.Location,
					Fqdn:          model.Properties.Fqdn,
				}

				if model.ExtendedLocation != nil {
					state.ExtendedLocation = []ExtendedLocationModel{
						{
							Name: *model.ExtendedLocation.Name,
							Type: *model.ExtendedLocation.Type,
						},
					}
				}

				if model.Kind != nil {
					state.Kind = *model.Kind
				}

				if props.Port != nil {
					state.Port = *props.Port
				}

				if props.Credentials != nil {
					state.Credential = []CredentialModel{
						{
							Username: *props.Credentials.Username,
							Password: *props.Credentials.Password,
						},
					}
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

func (r VcenterResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.ConnectedVmware.VcenterClient
			id, err := vcenters.ParseVCenterID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			metadata.Logger.Infof("deleting %s", *id)

			if resp, err := client.Delete(ctx, *id, vcenters.DefaultDeleteOperationOptions()); err != nil {
				if !response.WasNotFound(resp.HttpResponse) {
					return fmt.Errorf("deleting %s: %+v", *id, err)
				}
			}
			return nil
		},
	}
}

func (r VcenterResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return vcenters.ValidateVCenterID
}

func (r VcenterResource) Update() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.ConnectedVmware.VcenterClient
			id, err := vcenters.ParseVCenterID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			var state VcenterResourceModel
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

				patch := vcenters.ResourcePatch{
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
