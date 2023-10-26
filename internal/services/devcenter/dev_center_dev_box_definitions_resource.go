package devcenter

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonschema"
	"github.com/hashicorp/go-azure-sdk/resource-manager/devcenter/2023-04-01/devboxdefinitions"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/validation"
	"github.com/hashicorp/terraform-provider-azurerm/utils"
	"time"
)

var _ sdk.Resource = DevCenterDevBoxDefinitionsResource{}
var _ sdk.ResourceWithUpdate = DevCenterDevBoxDefinitionsResource{}

type DevCenterDevBoxDefinitionsResource struct{}

func (r DevCenterDevBoxDefinitionsResource) ModelObject() interface{} {
	return &DevCenterDevBoxDefinitionsResourceSchema{}
}

type DevCenterDevBoxDefinitionsResourceSchema struct {
	Location          string                `tfschema:"location"`
	Name              string                `tfschema:"name"`
	DevCenterName     string                `tfschema:"dev_center_name"`
	ResourceGroupName string                `tfschema:"resource_group_name"`
	HibernateSupport  string                `tfschema:"hibernate_support"`
	ImageReference    []ImageReferenceModel `tfschema:"image_reference"`
	OsStorageType     string                `tfschema:"os_storage_type"`
	Sku               []SkuModel            `tfschema:"sku"`
	Tags              map[string]string     `tfschema:"tags"`
}

type ImageReferenceModel struct {
	Id string `tfschema:"id"`
}

type SkuModel struct {
	Capacity int64  `tfschema:"capacity"`
	Family   string `tfschema:"family"`
	Name     string `tfschema:"name"`
	Size     string `tfschema:"size"`
	Tier     string `tfschema:"tier"`
}

func (r DevCenterDevBoxDefinitionsResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return devboxdefinitions.ValidateDevCenterDevBoxDefinitionID
}

func (r DevCenterDevBoxDefinitionsResource) ResourceType() string {
	return "azurerm_dev_center_dev_box_definitions"
}

func (r DevCenterDevBoxDefinitionsResource) Arguments() map[string]*schema.Schema {
	return map[string]*pluginsdk.Schema{
		"location": commonschema.Location(),

		"name": {
			ForceNew: true,
			Required: true,
			Type:     schema.TypeString,
		},

		"resource_group_name": commonschema.ResourceGroupName(),

		"dev_center_name": {
			ForceNew: true,
			Required: true,
			Type:     schema.TypeString,
		},

		"hibernate_support": {
			Optional: true,
			Type:     schema.TypeString,
			ValidateFunc: validation.StringInSlice([]string{
				string(devboxdefinitions.HibernateSupportDisabled),
				string(devboxdefinitions.HibernateSupportEnabled),
			}, false),
		},

		"image_reference": {
			Optional: true,
			Type:     schema.TypeList,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"id": {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},

		"os_storage_type": {
			Optional:     true,
			Type:         schema.TypeString,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"sku": {
			Optional: true,
			Type:     schema.TypeList,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Type:     schema.TypeString,
						Required: true,
					},

					"capacity": {
						Type:     schema.TypeInt,
						Optional: true,
					},

					"family": {
						Type:     schema.TypeString,
						Optional: true,
					},

					"size": {
						Type:     schema.TypeString,
						Optional: true,
					},

					"tier": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateFunc: validation.StringInSlice([]string{
							string(devboxdefinitions.SkuTierFree),
							string(devboxdefinitions.SkuTierBasic),
							string(devboxdefinitions.SkuTierPremium),
							string(devboxdefinitions.SkuTierStandard),
						}, false),
					},
				},
			},
		},

		"tags": commonschema.Tags(),
	}
}

func (r DevCenterDevBoxDefinitionsResource) Attributes() map[string]*schema.Schema {
	return map[string]*pluginsdk.Schema{}
}

func (r DevCenterDevBoxDefinitionsResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {

			var model DevCenterDevBoxDefinitionsResourceSchema
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			client := metadata.Client.DevCenter.V20230401.DevBoxDefinitions
			subscriptionId := metadata.Client.Account.SubscriptionId

			id := devboxdefinitions.NewDevCenterDevBoxDefinitionID(subscriptionId, model.ResourceGroupName, model.DevCenterName, model.Name)

			existing, err := client.Get(ctx, id)
			if err != nil {
				if !response.WasNotFound(existing.HttpResponse) {
					return fmt.Errorf("checking for the presence of an existing %s: %+v", id, err)
				}
			}
			if !response.WasNotFound(existing.HttpResponse) {
				return metadata.ResourceRequiresImport(r.ResourceType(), id)
			}

			sku, err := expandDevBoxDefinitionsSku(model.Sku)
			if err != nil {
				return fmt.Errorf("expanding sku: %+v", err)
			}

			imageReference, err := expandDevBoxDefinitionsImageReference(model.ImageReference)
			if err != nil {
				return fmt.Errorf("expanding image reference: %+v", err)
			}
			props := devboxdefinitions.DevBoxDefinitionProperties{
				HibernateSupport: pointer.To(devboxdefinitions.HibernateSupport(model.HibernateSupport)),
				ImageReference:   pointer.From(imageReference),
				OsStorageType:    pointer.To(model.OsStorageType),
				Sku:              pointer.From(sku),
			}

			definition := devboxdefinitions.DevBoxDefinition{
				Id:         utils.String(id.ID()),
				Location:   model.Location,
				Name:       utils.String(model.Name),
				Properties: pointer.To(props),
				Tags:       pointer.To(model.Tags),
			}

			if err := client.CreateOrUpdateThenPoll(ctx, id, definition); err != nil {
				return fmt.Errorf("creating %s: %+v", id, err)
			}

			metadata.SetID(id)
			return nil
		},
	}
}

func (r DevCenterDevBoxDefinitionsResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.DevCenter.V20230401.DevBoxDefinitions
			id, err := devboxdefinitions.ParseDevCenterDevBoxDefinitionID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			resp, err := client.Get(ctx, *id)
			if err != nil {
				if response.WasNotFound(resp.HttpResponse) {
					return metadata.MarkAsGone(*id)
				}
				return fmt.Errorf("reading %s: %+v", *id, err)
			}

			if model := resp.Model; model != nil {
				props := model.Properties
				state := DevCenterDevBoxDefinitionsResourceSchema{
					Location:          model.Location,
					Name:              *model.Name,
					ResourceGroupName: id.ResourceGroupName,
					DevCenterName:     id.DevCenterName,
					HibernateSupport:  string(pointer.From(props.HibernateSupport)),
					OsStorageType:     pointer.From(props.OsStorageType),
					Tags:              *model.Tags,
				}
				sku := flattenDevBoxDefinitionsSku(pointer.To(props.Sku))
				state.Sku = sku

				imageReference := flattenDevBoxDefinitionsImageReference(pointer.To(props.ImageReference))
				state.ImageReference = imageReference

				return metadata.Encode(&state)
			}
			return nil
		},
	}
}

func (r DevCenterDevBoxDefinitionsResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.DevCenter.V20230401.DevBoxDefinitions
			id, err := devboxdefinitions.ParseDevCenterDevBoxDefinitionID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			metadata.Logger.Infof("deleting %s", *id)

			if err := client.DeleteThenPoll(ctx, *id); err != nil {
				return fmt.Errorf("delete %s: %+v", *id, err)
			}

			return nil
		},
	}
}

func (r DevCenterDevBoxDefinitionsResource) Update() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.DevCenter.V20230401.DevBoxDefinitions
			id, err := devboxdefinitions.ParseDevCenterDevBoxDefinitionID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			var state DevCenterDevBoxDefinitionsResourceSchema
			if err := metadata.Decode(&state); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			sku, err := expandDevBoxDefinitionsSku(state.Sku)
			if err != nil {
				return fmt.Errorf("expanding sku: %+v", err)
			}
			imageReference, err := expandDevBoxDefinitionsImageReference(state.ImageReference)
			if err != nil {
				return fmt.Errorf("expanding image reference: %+v", err)
			}

			if metadata.ResourceData.HasChangesExcept("name", "resource_group_name", "location") {
				devBoxDefinitions := devboxdefinitions.DevBoxDefinitionUpdate{
					Properties: pointer.To(devboxdefinitions.DevBoxDefinitionUpdateProperties{
						HibernateSupport: pointer.To(devboxdefinitions.HibernateSupport(state.HibernateSupport)),
						ImageReference:   imageReference,
						OsStorageType:    pointer.To(state.OsStorageType),
						Sku:              sku,
					}),
					Tags: pointer.To(state.Tags),
				}
				if err := client.UpdateThenPoll(ctx, *id, devBoxDefinitions); err != nil {
					return fmt.Errorf("updating %s: %+v", *id, err)
				}
			}
			return nil
		},
	}
}
