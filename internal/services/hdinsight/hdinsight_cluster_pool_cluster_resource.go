package hdinsight

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/validation"
)

type ClusterPoolClusterResource struct{}

type ClusterPoolClusterModel struct {
	Name           string                  `json:"name"`
	ResourceGroup  string                  `json:"resource_group_name"`
	Location       string                  `json:"location"`
	ClusterProfile []ClusterProfile        `json:"cluster_profile"`
	ComputeProfile []ClusterComputeProfile `json:"compute_profile"`
	ClusterType    string                  `json:"cluster_type"`
	Tags           map[string]string       `json:"tags"`
}

type ClusterProfile struct {
	AuthorizationProfile []AuthorizationProfile `json:"authorization_profile"`
	AutoscaleProfile     []AutoscaleProfile     `json:"autoscale_profile"`
	FlinkProfile         []FlinkProfile         `json:"flink_profile"`
	IdentityProfile      []IdentityProfile      `json:"identity_profile"`
}

var _ sdk.ResourceWithUpdate = ClusterPoolClusterResource{}

func (c ClusterPoolClusterResource) ModelObject() interface{} {
	return &ClusterPoolClusterModel{}
}

func (c ClusterPoolClusterResource) ResourceType() string {
	return "azurerm_hdinsight_cluster_pool_cluster"
}

func (c ClusterPoolClusterResource) Arguments() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:         schema.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"resource_group_name": {
			Type:         schema.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"location": {
			Type:         schema.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"cluster_profile": {
			Type:     schema.TypeList,
			Required: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"authorization_profile": ClusterPoolAuthorizationProfileSchema(),
				},
			},
		},
	}
}

func (c ClusterPoolClusterResource) Attributes() map[string]*schema.Schema {
	//TODO implement me
	panic("implement me")
}

func (c ClusterPoolClusterResource) Create() sdk.ResourceFunc {
	//TODO implement me
	panic("implement me")
}

func (c ClusterPoolClusterResource) Read() sdk.ResourceFunc {
	//TODO implement me
	panic("implement me")
}

func (c ClusterPoolClusterResource) Delete() sdk.ResourceFunc {
	//TODO implement me
	panic("implement me")
}

func (c ClusterPoolClusterResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	//TODO implement me
	panic("implement me")
}

func (c ClusterPoolClusterResource) Update() sdk.ResourceFunc {
	//TODO implement me
	panic("implement me")
}
