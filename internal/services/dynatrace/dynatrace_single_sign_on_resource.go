package dynatrace

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonschema"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dynatrace/2023-04-27/singlesignon"
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/validation"
	"time"
)

type SingleSignOnResource struct{}

var _ sdk.ResourceWithUpdate = SingleSignOnResource{}

type SingleSignOnResourceModel struct {
	Name              string   `tfschema:"name"`
	ResourceGroup     string   `tfschema:"resource_group_name"`
	monitorName       string   `tfschema:"monitor_name"`
	AadDomains        []string `tfschema:"aad_domains"`
	EnterpriseAddId   string   `tfschema:"enterprise_app_id"`
	SingleSignOnState string   `tfschema:"single_sign_on_state"`
	SingleSignOnUrl   string   `tfschema:"single_sign_on_url"`
}

func (r SingleSignOnResource) Arguments() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"name": {
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"resource_group_name": commonschema.ResourceGroupName(),

		"monitor_name": {
			Type:         pluginsdk.TypeString,
			Required:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"aad_domains": {
			Type:     pluginsdk.TypeList,
			Required: true,
			Elem: &pluginsdk.Schema{
				Type: pluginsdk.TypeString,
			},
		},

		"enterprise_app_id": {
			Type:         pluginsdk.TypeString,
			Required:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"single_sign_on_state": {
			Type:     pluginsdk.TypeString,
			Required: true,
			ValidateFunc: validation.StringInSlice([]string{
				string(singlesignon.SingleSignOnStatesDisable),
				string(singlesignon.SingleSignOnStatesEnable),
				string(singlesignon.SingleSignOnStatesExisting),
				string(singlesignon.SingleSignOnStatesInitial),
			}, false),
		},

		"single_sign_on_url": {
			Type:         pluginsdk.TypeString,
			Required:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},
	}
}

func (r SingleSignOnResource) Attributes() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{}
}

func (r SingleSignOnResource) ModelObject() interface{} {
	return &SingleSignOnResourceModel{}
}

func (r SingleSignOnResource) ResourceType() string {
	return "azurerm_dynatrace_single_sign_on"
}

func (r SingleSignOnResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return singlesignon.ValidateSingleSignOnConfigurationID
}

func (r SingleSignOnResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Dynatrace.SingleSignOnClient
			subscriptionId := metadata.Client.Account.SubscriptionId

			var model SingleSignOnResourceModel

			if err := metadata.Decode(&model); err != nil {
				return err
			}

			id := singlesignon.NewSingleSignOnConfigurationID(subscriptionId, model.ResourceGroup, model.monitorName, model.Name)

			existing, err := client.Get(ctx, id)
			if err != nil && !response.WasNotFound(existing.HttpResponse) {
				return fmt.Errorf("checking for presence of %s: %+v", id, err)
			}
			if !response.WasNotFound(existing.HttpResponse) {
				return metadata.ResourceRequiresImport(r.ResourceType(), id)
			}

			props := singlesignon.DynatraceSingleSignOnProperties{
				AadDomains:        pointer.To(model.AadDomains),
				EnterpriseAppId:   pointer.To(model.EnterpriseAddId),
				SingleSignOnState: pointer.To(singlesignon.SingleSignOnStates(model.SingleSignOnState)),
				SingleSignOnURL:   pointer.To(model.SingleSignOnUrl),
			}

			singleSignOn := singlesignon.DynatraceSingleSignOnResource{
				Properties: props,
			}

			if err := client.CreateOrUpdateThenPoll(ctx, id, singleSignOn); err != nil {
				return fmt.Errorf("creating %s: %+v", id, err)
			}

			metadata.SetID(id)

			return nil
		},
	}
}

func (r SingleSignOnResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Dynatrace.SingleSignOnClient
			id, err := singlesignon.ParseSingleSignOnConfigurationID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			resp, err := client.Get(ctx, *id)
			if err != nil {
				if response.WasNotFound(resp.HttpResponse) {
					return metadata.MarkAsGone(id)
				}
				return fmt.Errorf("getting %s: %+v", id, err)
			}

			if model := resp.Model; model != nil {
				props := model.Properties

				state := SingleSignOnResourceModel{
					Name:              id.SingleSignOnConfigurationName,
					ResourceGroup:     id.ResourceGroupName,
					monitorName:       id.MonitorName,
					AadDomains:        pointer.From(props.AadDomains),
					EnterpriseAddId:   pointer.From(props.EnterpriseAppId),
					SingleSignOnState: string(pointer.From(props.SingleSignOnState)),
					SingleSignOnUrl:   pointer.From(props.SingleSignOnURL),
				}

				return metadata.Encode(&state)
			}
			return nil
		},
	}
}

func (r SingleSignOnResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Dynatrace.SingleSignOnClient
			id, err := singlesignon.ParseSingleSignOnConfigurationID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			metadata.Logger.Infof("deleting %s", *id)

			props := singlesignon.DynatraceSingleSignOnResource{}

			if _, err := client.CreateOrUpdate(ctx, *id, props); err != nil {
				return fmt.Errorf("deleting %s: %+v", id, err)
			}

			return nil
		},
	}
}

func (r SingleSignOnResource) Update() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Dynatrace.SingleSignOnClient
			id, err := singlesignon.ParseSingleSignOnConfigurationID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			var state SingleSignOnResourceModel
			if err := metadata.Decode(&state); err != nil {
				return err
			}

			existing, err := client.Get(ctx, *id)
			if err != nil {
				return err
			}
			model := existing.Model

			if metadata.ResourceData.HasChange("aad_domains") {
				model.Properties.AadDomains = pointer.To(state.AadDomains)
			}

			if metadata.ResourceData.HasChange("enterprise_app_id") {
				model.Properties.EnterpriseAppId = pointer.To(state.EnterpriseAddId)
			}

			if metadata.ResourceData.HasChange("single_sign_on_state") {
				model.Properties.SingleSignOnState = pointer.To(singlesignon.SingleSignOnStates(state.SingleSignOnState))
			}

			if metadata.ResourceData.HasChange("single_sign_on_url") {
				model.Properties.SingleSignOnURL = pointer.To(state.SingleSignOnUrl)
			}

			if _, err := client.CreateOrUpdate(ctx, *id, *model); err != nil {
				return fmt.Errorf("updating %s: %+v", id, err)
			}

			return nil
		},
	}
}
