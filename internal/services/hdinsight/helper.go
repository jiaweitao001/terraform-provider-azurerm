package hdinsight

import (
	"github.com/hashicorp/go-azure-sdk/resource-manager/hdinsight/2024-05-01/hdinsights"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/validation"
)

type AuthorizationProfile struct {
	GroupIds []string `json:"group_ids"`
	UserIds  []string `json:"user_ids"`
}

type AutoscaleProfile struct {
	AutoscaleType               string                `json:"autoscale_type"`
	AutoscaleEnabled            bool                  `json:"autoscale_enabled"`
	GracefulDecommissionTimeout int                   `json:"graceful_decommission_timeout"`
	LoadBasedConfig             []LoadBasedConfig     `json:"load_based_config"`
	ScheduleBasedConfig         []ScheduleBasedConfig `json:"schedule_based_config"`
}

type LoadBasedConfig struct {
	CooldownPeriod string        `json:"cooldown_period"`
	MaxNodes       int           `json:"max_nodes"`
	MinNodes       int           `json:"min_nodes"`
	PollInterval   string        `json:"poll_interval"`
	ScalingRules   []ScalingRule `json:"scaling_rules"`
}

type ScalingRule struct {
	ActionType      string           `json:"action_type"`
	EvaluationCount int              `json:"evaluation_count"`
	ScalingMetric   string           `json:"scaling_metric"`
	ComparisonRule  []ComparisonRule `json:"comparison_rule"`
}

type ComparisonRule struct {
	Operator  string `json:"operator"`
	Threshold int    `json:"threshold"`
}

type ScheduleBasedConfig struct {
	DefaultCount int        `json:"default_count"`
	Schedules    []Schedule `json:"schedule"`
	TimeZone     string     `json:"time_zone"`
}

type Schedule struct {
	Count     int    `json:"count"`
	Days      string `json:"days"`
	EndTime   string `json:"end_time"`
	StartTime string `json:"start_time"`
}

type FlinkProfile struct {
	DeploymentMode string          `json:"deployment_mode"`
	NumReplicas    int             `json:"num_replicas"`
	CatalogOptions []CatalogOption `json:"catalog_option"`
	HistoryServer  []HistoryServer `json:"history_server"`
	JobManager     []JobManager    `json:"job_manager"`
	JobSpec        []JobSpec       `json:"job_spec"`
	Storage        []Storage       `json:"storage"`
	TaskManager    []TaskManager   `json:"task_manager"`
}

type CatalogOption struct {
	MetastoreDbConnectionAuthenticationMode string `json:"metastore_db_connection_authentication_mode"`
	MetastoreDbConnectionPassword           string `json:"metastore_db_connection_password"`
	MetastoreDbConnectionUri                string `json:"metastore_db_connection_uri"`
	MetastoreDbConnectionUserName           string `json:"metastore_db_connection_user_name"`
}

type HistoryServer struct {
	Cpu    int `json:"cpu"`
	Memory int `json:"memory"`
}

type JobManager struct {
	Cpu    int `json:"cpu"`
	Memory int `json:"memory"`
}

type JobSpec struct {
	Args            string `json:"args"`
	EntryClass      string `json:"entry_class"`
	JarName         string `json:"jar_name"`
	JobJarDirectory string `json:"job_jar_directory"`
	UpgradeMode     string `json:"upgrade_mode"`
	SavePointName   string `json:"save_point_name"`
}

type Storage struct {
	StorageUrl string `json:"storage_url"`
	StorageKey string `json:"storage_key"`
}

type TaskManager struct {
	Cpu    int `json:"cpu"`
	Memory int `json:"memory"`
}

type IdentityProfile struct {
	MsiClientId   string `json:"msi_client_id"`
	MsiObjectId   string `json:"msi_object_id"`
	MsiResourceId string `json:"msi_resource_id"`
}

type ClusterComputeProfile struct {
	NodeCount int    `json:"node_count"`
	NodeType  string `json:"node_type"`
	VmSize    string `json:"vm_size"`
}

func ClusterPoolAuthorizationProfileSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Required: true,
		MaxItems: 1,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"authorization_profile": ClusterAuthorizationProfileSchema(),

				"autoscale_profile": ClusterAutoscaleProfileSchema(),

				"flink_profile": ClusterFlinkProfileSchema(),

				"identity_profile": ClusterIdentityProfileSchema(),
			},
		},
	}
}

func ClusterAuthorizationProfileSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type: pluginsdk.TypeList,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"autoscale_enabled": {
					Type:     pluginsdk.TypeBool,
					Required: true,
				},

				"autoscale_type": {
					Type:     pluginsdk.TypeString,
					Optional: true,
					ValidateFunc: validation.StringInSlice([]string{
						string(hdinsights.AutoscaleTypeLoadBased),
						string(hdinsights.AutoscaleTypeScheduleBased),
					}, false),
				},

				"graceful_decommission_timeout": {
					Type:     pluginsdk.TypeInt,
					Optional: true,
				},

				"load_based_config": ClusterLoadBasedConfigSchema(),

				"schedule_based_config": ClusterScheduleBasedConfigSchema(),
			},
		},
	}
}

func ClusterLoadBasedConfigSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type: pluginsdk.TypeList,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"max_nodes": {
					Type:     pluginsdk.TypeInt,
					Required: true,
				},

				"min_nodes": {
					Type:     pluginsdk.TypeInt,
					Required: true,
				},

				"scaling_rules": ClusterScalingRuleSchema(),

				"poll_interval": {
					Type:     pluginsdk.TypeString,
					Optional: true,
				},

				"cooldown_period": {
					Type:     pluginsdk.TypeString,
					Optional: true,
				},
			},
		},
	}
}

func ClusterScalingRuleSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type: pluginsdk.TypeList,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"action_type": {
					Type:     pluginsdk.TypeString,
					Required: true,
				},

				"evaluation_count": {
					Type:     pluginsdk.TypeInt,
					Required: true,
				},

				"scaling_metric": {
					Type:     pluginsdk.TypeString,
					Required: true,
				},

				"comparison_rule": ClusterComparisonRuleSchema(),
			},
		},
	}
}

func ClusterComparisonRuleSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type: pluginsdk.TypeList,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"operator": {
					Type:     pluginsdk.TypeString,
					Required: true,
				},

				"threshold": {
					Type:     pluginsdk.TypeInt,
					Required: true,
				},
			},
		},
	}
}

func ClusterScheduleBasedConfigSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type: pluginsdk.TypeList,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"default_count": {
					Type:     pluginsdk.TypeInt,
					Required: true,
				},

				"schedules": ClusterScheduleSchema(),

				"time_zone": {
					Type:     pluginsdk.TypeString,
					Required: true,
				},
			},
		},
	}
}

func ClusterScheduleSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type: pluginsdk.TypeList,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"count": {
					Type:     pluginsdk.TypeInt,
					Required: true,
				},

				"days": {
					Type:     pluginsdk.TypeString,
					Required: true,
					ValidateFunc: validation.StringInSlice([]string{
						string(hdinsights.ScheduleDayMonday),
						string(hdinsights.ScheduleDayTuesday),
						string(hdinsights.ScheduleDayWednesday),
						string(hdinsights.ScheduleDayThursday),
						string(hdinsights.ScheduleDayFriday),
						string(hdinsights.ScheduleDaySaturday),
						string(hdinsights.ScheduleDaySunday),
					}, false),
				},

				"end_time": {
					Type:     pluginsdk.TypeString,
					Required: true,
				},

				"start_time": {
					Type:     pluginsdk.TypeString,
					Required: true,
				},
			},
		},
	}
}
