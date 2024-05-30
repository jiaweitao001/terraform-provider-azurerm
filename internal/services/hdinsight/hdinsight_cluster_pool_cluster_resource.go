package hdinsight

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/location"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hdinsight/2024-05-01/hdinsights"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tags"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/validation"
	"time"
)

type ClusterPoolClusterResource struct{}

type ClusterPoolClusterModel struct {
	Name            string                  `tfschema:"name"`
	ClusterPoolName string                  `tfschema:"cluster_pool_name"`
	ResourceGroup   string                  `tfschema:"resource_group_name"`
	Location        string                  `tfschema:"location"`
	ClusterProfile  []ClusterProfile        `tfschema:"cluster_profile"`
	ComputeProfile  []ClusterComputeProfile `tfschema:"compute_profile"`
	ClusterType     string                  `tfschema:"cluster_type"`
	Tags            map[string]string       `tfschema:"tags"`
}

type ClusterProfile struct {
	AuthorizationProfile   []AuthorizationProfile       `tfschema:"authorization_profile"`
	AutoscaleProfile       []AutoscaleProfile           `tfschema:"autoscale_profile"`
	FlinkProfile           []FlinkProfile               `tfschema:"flink_profile"`
	IdentityProfile        []IdentityProfile            `tfschema:"identity_profile"`
	KafkaProfile           []KafkaProfile               `tfschema:"kafka_profile"`
	LogAnalyticsProfile    []ClusterLogAnalyticsProfile `tfschema:"log_analytics_profile"`
	ManagedIdentityProfile []ManagedIdentityProfile     `tfschema:"managed_identity_profile"`
	ClusterAccessProfile   []ClusterAccessProfile       `tfschema:"cluster_access_profile"`
	PrometheusProfile      []PrometheusProfile          `tfschema:"prometheus_profile"`
	RangerProfile          []RangerProfile              `tfschema:"ranger_profile"`
	ScriptActionProfile    []ScriptActionProfile        `tfschema:"script_action_profile"`
	SecretsProfile         []SecretsProfile             `tfschema:"secrets_profile"`
	ServiceConfigsProfiles []ServiceConfigsProfile      `tfschema:"service_configs_profiles"`
	SparkProfile           []SparkProfile               `tfschema:"spark_profile"`
	SshProfile             []SshProfile                 `tfschema:"ssh_profile"`
	TrinoProfile           []TrinoProfile               `tfschema:"trino_profile"`
	ClusterType            string                       `tfschema:"cluster_type"`
	ClusterVersion         string                       `tfschema:"cluster_version"`
	OssVersion             string                       `tfschema:"oss_version"`
}

var _ sdk.ResourceWithUpdate = ClusterPoolClusterResource{}

func (r ClusterPoolClusterResource) ModelObject() interface{} {
	return &ClusterPoolClusterModel{}
}

func (r ClusterPoolClusterResource) ResourceType() string {
	return "azurerm_hdinsight_cluster_pool_cluster"
}

func (r ClusterPoolClusterResource) Arguments() map[string]*schema.Schema {
	return map[string]*pluginsdk.Schema{
		"name": {
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"cluster_pool_name": {
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"resource_group_name": {
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"location": {
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"cluster_profile": ClusterProfileSchema(),

		"cluster_type": {
			Type:         pluginsdk.TypeString,
			Required:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"compute_profile": ClusterComputeProfileSchema(),

		"tags": tags.Schema(),
	}
}

func (r ClusterPoolClusterResource) Attributes() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func (r ClusterPoolClusterResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			var model ClusterPoolClusterModel
			if err := metadata.Decode(&model); err != nil {
				return err
			}

			client := metadata.Client.HDInsight2024.Hdinsights
			subscriptionId := metadata.Client.Account.SubscriptionId

			if err := metadata.Decode(&model); err != nil {
				return err
			}
			id := hdinsights.NewClusterID(subscriptionId, model.ResourceGroup, model.ClusterPoolName, model.Name)

			existing, err := client.ClustersGet(ctx, id)
			if err != nil {
				if !response.WasNotFound(existing.HttpResponse) {
					return fmt.Errorf("checking for presence of existing %s: %+v", id, err)
				}
			}

			if !response.WasNotFound(existing.HttpResponse) {
				return metadata.ResourceRequiresImport(r.ResourceType(), id)
			}

			cluster := hdinsights.Cluster{
				Location: model.Location,
				Properties: &hdinsights.ClusterResourceProperties{
					ClusterType:    model.ClusterType,
					ClusterProfile: expandClusterProfile(model.ClusterProfile),
					ComputeProfile: pointer.From(expandClusterComputeProfile(model.ComputeProfile)),
				},
				Tags: pointer.To(model.Tags),
			}

			if err := client.ClustersCreateThenPoll(ctx, id, cluster); err != nil {
				return fmt.Errorf("creating %s: %+v", id, err)
			}

			metadata.SetID(id)
			return nil
		},
	}
}

func (r ClusterPoolClusterResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.HDInsight2024.Hdinsights
			id, err := hdinsights.ParseClusterID(metadata.ResourceData.Id())
			if err != nil {
				return fmt.Errorf("parsing %s: %+v", metadata.ResourceData.Id(), err)
			}

			resp, err := client.ClustersGet(ctx, *id)
			if err != nil {
				if response.WasNotFound(resp.HttpResponse) {
					return metadata.MarkAsGone(id)
				}
				return fmt.Errorf("reading %s: %+v", *id, err)
			}

			var state ClusterPoolClusterModel
			state.Name = id.ClusterName
			state.ResourceGroup = id.ResourceGroupName

			if model := resp.Model; model != nil {
				state.Location = location.Normalize(model.Location)
				state.Tags = pointer.From(model.Tags)

				if props := model.Properties; props != nil {
					state.ClusterType = props.ClusterType
					state.ClusterProfile = flattenClusterProfile(props.ClusterProfile)
					state.ComputeProfile = flattenClusterComputeProfile(props.ComputeProfile)
				}
			}

			return metadata.Encode(&state)
		},
	}
}

func (r ClusterPoolClusterResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.HDInsight2024.Hdinsights
			id, err := hdinsights.ParseClusterID(metadata.ResourceData.Id())
			if err != nil {
				return fmt.Errorf("parsing %s: %+v", metadata.ResourceData.Id(), err)
			}

			if err := client.ClustersDeleteThenPoll(ctx, *id); err != nil {
				return fmt.Errorf("deleting %s: %+v", *id, err)
			}

			return nil
		},
	}
}

func (r ClusterPoolClusterResource) Update() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			var state ClusterPoolClusterModel
			if err := metadata.Decode(&state); err != nil {
				return err
			}

			client := metadata.Client.HDInsight2024.Hdinsights
			id, err := hdinsights.ParseClusterID(metadata.ResourceData.Id())
			if err != nil {
				return fmt.Errorf("parsing %s: %+v", metadata.ResourceData.Id(), err)
			}

			cluster := hdinsights.ClusterPatch{
				Properties: &hdinsights.ClusterPatchProperties{
					ClusterProfile: expandClusterProfilePatch(state.ClusterProfile),
				},
				Tags: pointer.To(state.Tags),
			}

			if err := client.ClustersUpdateThenPoll(ctx, *id, cluster); err != nil {
				return fmt.Errorf("updating %s: %+v", *id, err)
			}

			return nil
		},
	}
}

func (r ClusterPoolClusterResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return hdinsights.ValidateClusterID
}
