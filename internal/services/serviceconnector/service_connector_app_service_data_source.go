package serviceconnector

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/appservice/validate"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/serviceconnector/sdk/2022-05-01/servicelinker"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"time"
)

type AppServiceConnectorDataSource struct{}

type AppServiceConnectorDataSourceModel struct {
	Name             string          `tfschema:"name"`
	ResourceGroup    string          `tfschema:"resource_group_name"`
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
			ValidateFunc: validate.WebAppName,
		},
	}
}

func (r AppServiceConnectorDataSource) Attributes() map[string]*schema.Schema {
	//TODO implement me
	panic("implement me")
}

func (r AppServiceConnectorDataSource) ModelObject() interface{} {
	return &AppServiceConnectorDataSourceModel{}
}

func (r AppServiceConnectorDataSource) ResourceType() string {
	//TODO implement me
	panic("implement me")
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

			if !response.WasNotFound(existing.HttpResponse) {
				return metadata.ResourceRequiresImport(r.ResourceType(), id)
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
					TargetResourceId: flattenTargetService(props),
					AuthInfo:         flattenServiceConnectorAuthInfo(props.AuthInfo),
				}

				if props.ClientType != nil {
					state.ClientType = string(*props.ClientType)
				}

				if props.VNetSolution != nil && props.VNetSolution.Type != nil {
					state.VnetSolution = string(*props.VNetSolution.Type)
				}

				return metadata.Encode(&state)
			}
			return nil
		},
	}
}
