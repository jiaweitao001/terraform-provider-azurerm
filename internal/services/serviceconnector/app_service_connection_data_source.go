// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceconnector

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicelinker/2024-04-01/servicelinker"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/web/validate"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/validation"
)

type AppServiceConnectorDataSource struct{}

type AppServiceConnectorDataSourceModel struct {
	Name             string             `tfschema:"name"`
	AppServiceId     string             `tfschema:"app_service_id"`
	TargetResourceId string             `tfschema:"target_resource_id"`
	ClientType       string             `tfschema:"client_type"`
	AuthInfo         []AuthInfoModel    `tfschema:"authentication"`
	VnetSolution     string             `tfschema:"vnet_solution"`
	SecretStore      []SecretStoreModel `tfschema:"secret_store"`
}

var _ sdk.DataSource = AppServiceConnectorDataSource{}

func (r AppServiceConnectorDataSource) ResourceType() string {
	return "azurerm_app_service_connection"
}

func (r AppServiceConnectorDataSource) ModelObject() interface{} {
	return &AppServiceConnectorDataSourceModel{}
}

func (r AppServiceConnectorDataSource) Arguments() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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

		"vnet_solution": {
			Type:     pluginsdk.TypeString,
			Computed: true,
		},

		"secret_store": {
			Type:     pluginsdk.TypeList,
			Computed: true,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*schema.Schema{
					"key_vault_id": {
						Type:     pluginsdk.TypeString,
						Computed: true,
					},
				},
			},
		},

		"authentication": {
			Type:     pluginsdk.TypeList,
			Computed: true,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*schema.Schema{
					"type": {
						Type:     pluginsdk.TypeString,
						Computed: true,
					},

					"name": {
						Type:     pluginsdk.TypeString,
						Computed: true,
					},

					"secret": {
						Type:      pluginsdk.TypeString,
						Computed:  true,
						Sensitive: true,
					},

					"client_id": {
						Type:     pluginsdk.TypeString,
						Computed: true,
					},

					"subscription_id": {
						Type:     pluginsdk.TypeString,
						Computed: true,
					},

					"principal_id": {
						Type:     pluginsdk.TypeString,
						Computed: true,
					},

					"certificate": {
						Type:      pluginsdk.TypeString,
						Computed:  true,
						Sensitive: true,
					},
				},
			},
		},
	}
}

func (r AppServiceConnectorDataSource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.ServiceConnector.ServiceLinkerClient

			var config AppServiceConnectorDataSourceModel
			if err := metadata.Decode(&config); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			id := servicelinker.NewScopedLinkerID(config.AppServiceId, config.Name)

			resp, err := client.LinkerGet(ctx, id)
			if err != nil {
				if response.WasNotFound(resp.HttpResponse) {
					return fmt.Errorf("App Service Connection %q was not found in App Service %q", config.Name, config.AppServiceId)
				}
				return fmt.Errorf("retrieving %s: %+v", id, err)
			}

			if model := resp.Model; model != nil {
				props := model.Properties
				if props.AuthInfo == nil || props.TargetService == nil {
					return fmt.Errorf("retrieving %s: properties were nil", id)
				}

				state := AppServiceConnectorDataSourceModel{
					Name:             id.LinkerName,
					AppServiceId:     id.ResourceUri,
					TargetResourceId: flattenTargetService(props.TargetService),
					AuthInfo:         flattenServiceConnectorAuthInfoForDataSource(props.AuthInfo),
				}

				if props.ClientType != nil {
					state.ClientType = string(*props.ClientType)
				}

				if props.VNetSolution != nil && props.VNetSolution.Type != nil {
					state.VnetSolution = string(*props.VNetSolution.Type)
				}

				if props.SecretStore != nil {
					state.SecretStore = flattenSecretStore(*props.SecretStore)
				}

				return metadata.Encode(&state)
			}

			return fmt.Errorf("retrieving %s: model was nil", id)
		},
	}
}

func flattenServiceConnectorAuthInfoForDataSource(input servicelinker.AuthInfoBase) []AuthInfoModel {
	var authType string
	var name string
	var clientId string
	var principalId string
	var subscriptionId string

	if value, ok := input.(servicelinker.SecretAuthInfo); ok {
		authType = string(servicelinker.AuthTypeSecret)
		if value.Name != nil {
			name = *value.Name
		}
	}

	if _, ok := input.(servicelinker.SystemAssignedIdentityAuthInfo); ok {
		authType = string(servicelinker.AuthTypeSystemAssignedIdentity)
	}

	if value, ok := input.(servicelinker.UserAssignedIdentityAuthInfo); ok {
		authType = string(servicelinker.AuthTypeUserAssignedIdentity)
		if value.ClientId != nil {
			clientId = *value.ClientId
		}
		if value.SubscriptionId != nil {
			subscriptionId = *value.SubscriptionId
		}
	}

	if value, ok := input.(servicelinker.ServicePrincipalSecretAuthInfo); ok {
		authType = string(servicelinker.AuthTypeServicePrincipalSecret)
		clientId = value.ClientId
		principalId = value.PrincipalId
	}

	if value, ok := input.(servicelinker.ServicePrincipalCertificateAuthInfo); ok {
		authType = string(servicelinker.AuthTypeServicePrincipalCertificate)
		clientId = value.ClientId
		principalId = value.PrincipalId
	}

	return []AuthInfoModel{
		{
			Type:           authType,
			Name:           name,
			ClientId:       clientId,
			PrincipalId:    principalId,
			SubscriptionId: subscriptionId,
		},
	}
}
