package serviceconnector

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/serviceconnector/sdk/2022-05-01/servicelinker"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/web/validate"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/validation"
)

type AppServiceConnectorDataSource struct{}

type AppServiceConnectorDataSourceModel struct {
	Name             string          `tfschema:"name"`
	AppServiceId     string          `tfschema:"app_service_id"`
	TargetResourceId string          `tfschema:"target_resource_id"`
	ClientType       string          `tfschema:"client_type"`
	AuthInfo         []AuthInfoModel `tfschema:"auth_info"`
	VnetSolution     string          `tfschema:"vnet_solution"`
}

var _ sdk.DataSource = AppServiceConnectorDataSource{}

func (r AppServiceConnectorDataSource) Arguments() map[string]*schema.Schema {
	return map[string]*pluginsdk.Schema{
		"name": {
			Type:         pluginsdk.TypeString,
			Required:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"app_service_id": {
			Type:         pluginsdk.TypeString,
			Required:     true,
			ValidateFunc: validate.AppServiceID,
		},
	}
}

func (r AppServiceConnectorDataSource) Attributes() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"target_resource_id": {
			Type:     pluginsdk.TypeString,
			Computed: true,
		},

		"client_type": {
			Type:     pluginsdk.TypeString,
			Computed: true,
		},

		"auth_info": authInfoSchemaComputed(),

		"vnet_solution": {
			Type:     pluginsdk.TypeString,
			Computed: true,
		},
	}
}

func (r AppServiceConnectorDataSource) ModelObject() interface{} {
	return &AppServiceConnectorDataSourceModel{}
}

func (r AppServiceConnectorDataSource) ResourceType() string {
	return "azurerm_app_service_connection"
}

func (r AppServiceConnectorDataSource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.ServiceConnector.ServiceLinkerClient

			var serviceConnector AppServiceConnectorDataSourceModel
			if err := metadata.Decode(&serviceConnector); err != nil {
				return err
			}

			id := servicelinker.NewScopedLinkerID(serviceConnector.AppServiceId, serviceConnector.Name)

			existing, err := client.LinkerGet(ctx, id)
			if err != nil {
				if response.WasNotFound(existing.HttpResponse) {
					return metadata.MarkAsGone(id)
				}
				return fmt.Errorf("reading %s: %+v", id, err)
			}

			if model := existing.Model; model != nil {
				props := model.Properties
				if props.AuthInfo == nil {
					return fmt.Errorf("reading authentication info for service connector %s", id)
				}
				if props.TargetService == nil {
					return fmt.Errorf("reading target service for service connector %s", id)
				}

				state := AppServiceConnectorDataSourceModel{
					Name:             id.LinkerName,
					AppServiceId:     id.ResourceUri,
					TargetResourceId: flattenTargetService(props.TargetService),
					AuthInfo:         flattenServiceConnectorAuthInfo(props.AuthInfo),
				}

				if props.ClientType != nil {
					state.ClientType = string(*props.ClientType)
				}

				if props.VNetSolution != nil && props.VNetSolution.Type != nil {
					state.VnetSolution = string(*props.VNetSolution.Type)
				}

				metadata.SetID(id)

				if err := metadata.Encode(&state); err != nil {
					return fmt.Errorf("encoding: %+v", err)
				}
			}
			return nil
		},
	}
}
