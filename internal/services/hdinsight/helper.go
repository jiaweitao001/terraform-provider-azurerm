package hdinsight

import (
	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hdinsight/2024-05-01/hdinsights"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/validation"
)

type AuthorizationProfile struct {
	GroupIds []string `tfschema:"group_ids"`
	UserIds  []string `tfschema:"user_ids"`
}

type AutoscaleProfile struct {
	AutoscaleType               string                `tfschema:"autoscale_type"`
	AutoscaleEnabled            bool                  `tfschema:"autoscale_enabled"`
	GracefulDecommissionTimeout int                   `tfschema:"graceful_decommission_timeout"`
	LoadBasedConfig             []LoadBasedConfig     `tfschema:"load_based_config"`
	ScheduleBasedConfig         []ScheduleBasedConfig `tfschema:"schedule_based_config"`
}

type KafkaProfile struct {
	DiskStorage           []DiskStorage `tfschema:"disk_storage"`
	KRaftEnabled          bool          `tfschema:"kraft_enabled"`
	PublicEndpointEnabled bool          `tfschema:"public_endpoint_enabled"`
	RemoteStorageUri      string        `tfschema:"remote_storage_uri"`
}

type DiskStorage struct {
	DiskSize int    `tfschema:"disk_size"`
	DiskType string `tfschema:"disk_type"`
}

type ClusterLogAnalyticsProfile struct {
	ApplicationLogs []ApplicationLogs `tfschema:"application_logs"`
	Enabled         bool              `tfschema:"enabled"`
	MetricsEnabled  bool              `tfschema:"metrics_enabled"`
}

type ApplicationLogs struct {
	StdErrorEnabled bool `tfschema:"std_error_enabled"`
	StdOutEnabled   bool `tfschema:"std_out_enabled"`
}

type LoadBasedConfig struct {
	CooldownPeriod int           `tfschema:"cooldown_period"`
	MaxNodes       int           `tfschema:"max_nodes"`
	MinNodes       int           `tfschema:"min_nodes"`
	PollInterval   int           `tfschema:"poll_interval"`
	ScalingRules   []ScalingRule `tfschema:"scaling_rules"`
}

type ScalingRule struct {
	ActionType      string           `tfschema:"action_type"`
	EvaluationCount int              `tfschema:"evaluation_count"`
	ScalingMetric   string           `tfschema:"scaling_metric"`
	ComparisonRule  []ComparisonRule `tfschema:"comparison_rule"`
}

type ComparisonRule struct {
	Operator  string  `tfschema:"operator"`
	Threshold float64 `tfschema:"threshold"`
}

type ScheduleBasedConfig struct {
	DefaultCount int        `tfschema:"default_count"`
	Schedules    []Schedule `tfschema:"schedule"`
	TimeZone     string     `tfschema:"time_zone"`
}

type Schedule struct {
	Count     int      `tfschema:"count"`
	Days      []string `tfschema:"days"`
	EndTime   string   `tfschema:"end_time"`
	StartTime string   `tfschema:"start_time"`
}

type FlinkProfile struct {
	DeploymentMode string               `tfschema:"deployment_mode"`
	NumReplicas    int                  `tfschema:"num_replicas"`
	CatalogOptions []FlinkCatalogOption `tfschema:"catalog_option"`
	HistoryServer  []ComputeResource    `tfschema:"history_server"`
	JobManager     []ComputeResource    `tfschema:"job_manager"`
	JobSpec        []JobSpec            `tfschema:"job_spec"`
	Storage        []Storage            `tfschema:"storage"`
	TaskManager    []ComputeResource    `tfschema:"task_manager"`
}

type FlinkCatalogOption struct {
	MetastoreDbConnectionAuthenticationMode string `tfschema:"metastore_db_connection_authentication_mode"`
	MetastoreDbConnectionPassword           string `tfschema:"metastore_db_connection_password"`
	MetastoreDbConnectionUrl                string `tfschema:"metastore_db_connection_url"`
	MetastoreDbConnectionUserName           string `tfschema:"metastore_db_connection_user_name"`
}

type ComputeResource struct {
	Cpu    float64 `tfschema:"cpu"`
	Memory int     `tfschema:"memory"`
}

type JobSpec struct {
	Args            string `tfschema:"args"`
	EntryClass      string `tfschema:"entry_class"`
	JarName         string `tfschema:"jar_name"`
	JobJarDirectory string `tfschema:"job_jar_directory"`
	UpgradeMode     string `tfschema:"upgrade_mode"`
	SavePointName   string `tfschema:"save_point_name"`
}

type Storage struct {
	StorageUrl string `tfschema:"storage_url"`
	StorageKey string `tfschema:"storage_key"`
}

type IdentityProfile struct {
	MsiClientId   string `tfschema:"msi_client_id"`
	MsiObjectId   string `tfschema:"msi_object_id"`
	MsiResourceId string `tfschema:"msi_resource_id"`
}

type ClusterComputeProfile struct {
	Node []Node `tfschema:"node"`
}

type Node struct {
	Count  int    `tfschema:"count"`
	Type   string `tfschema:"type"`
	VmSize string `tfschema:"vm_size"`
}

type ManagedIdentityProfile struct {
	Identities []Identity `tfschema:"identities"`
}

type ClusterAccessProfile struct {
	InternalIngressEnabled bool `tfschema:"internal_ingress_enabled"`
}

type Identity struct {
	ClientId   string `tfschema:"client_id"`
	ObjectId   string `tfschema:"object_id"`
	ResourceId string `tfschema:"resource_id"`
	Type       string `tfschema:"type"`
}

type PrometheusProfile struct {
	Enabled bool `tfschema:"enabled"`
}

type RangerProfile struct {
	RangerAdmin    []RangerAdmin    `tfschema:"ranger_admin"`
	StorageAccount string           `tfschema:"storage_account"`
	RangerUserSync []RangerUserSync `tfschema:"ranger_user_sync"`
}

type RangerAdmin struct {
	Admins   []string   `tfschema:"admins"`
	Database []Database `tfschema:"database"`
}

type Database struct {
	Host              string `tfschema:"host"`
	Name              string `tfschema:"name"`
	PasswordSecretRef string `tfschema:"password_secret_ref"`
	Username          string `tfschema:"username"`
}

type RangerUserSync struct {
	Enabled             bool     `tfschema:"enabled"`
	Groups              []string `tfschema:"groups"`
	Mode                string   `tfschema:"mode"` //rangerusersyncmode
	UserMappingLocation string   `tfschema:"user_mapping_location"`
	Users               []string `tfschema:"users"`
}

type ScriptActionProfile struct {
	Name             string   `tfschema:"name"`
	Services         []string `tfschema:"services"`
	Type             string   `tfschema:"type"`
	Url              string   `tfschema:"url"`
	Parameters       string   `tfschema:"parameters"`
	ShouldPersist    bool     `tfschema:"should_persist"`
	TimeoutInMinutes int      `tfschema:"timeout_in_minutes"`
}

type SecretsProfile struct {
	KeyVaultResourceId string   `tfschema:"key_vault_resource_id"`
	Secrets            []Secret `tfschema:"secrets"`
}

type Secret struct {
	KeyVaultObjectName string `tfschema:"key_vault_object_name"`
	Type               string `tfschema:"type"`
	ReferenceName      string `tfschema:"reference_name"`
	Version            string `tfschema:"version"`
}

type ServiceConfigsProfile struct {
	ServiceName string   `tfschema:"service_name"`
	Configs     []Config `tfschema:"configs"`
}

type Config struct {
	Component string `tfschema:"component"`
	Files     []File `tfschema:"files"`
}

type File struct {
	FileName string            `tfschema:"file_name"`
	Content  string            `tfschema:"content"`
	Encoding string            `tfschema:"encoding"`
	Path     string            `tfschema:"path"`
	Values   map[string]string `tfschema:"values"`
}

type SparkProfile struct {
	DefaultStorageUrl string          `tfschema:"default_storage_url"`
	MetastoreSpec     []MetastoreSpec `tfschema:"metastore_spec"`
}

type MetastoreSpec struct {
	DbConnectionAuthenticationMode string `tfschema:"db_connection_authentication_mode"`
	DbName                         string `tfschema:"db_name"`
	DbPasswordSecretName           string `tfschema:"db_password_secret_name"`
	DbServerHost                   string `tfschema:"db_server_host"`
	DbUserName                     string `tfschema:"db_user_name"`
	KeyVaultId                     string `tfschema:"key_vault_id"`
	ThriftUrl                      string `tfschema:"thrift_url"`
}

type SshProfile struct {
	Count  int    `tfschema:"count"`
	VmSize string `tfschema:"vm_size"`
}

type TrinoProfile struct {
	CatalogOptions    []TrinoCatalogOption `tfschema:"catalog_options"`
	Coordinator       []Coordinator        `tfschema:"coordinator"`
	Worker            []Worker             `tfschema:"worker"`
	UserPluginsSpec   []UserPluginsSpec    `tfschema:"user_plugins_spec"`
	UserTelemetrySpec []UserTelemetrySpec  `tfschema:"user_telemetry_spec"`
}

type TrinoCatalogOption struct {
	Hive []Hive `tfschema:"hive"`
}

type Hive struct {
	CatalogName                             string `tfschema:"catalog_name"`
	MetastoreDbConnectionAuthenticationMode string `tfschema:"metastore_db_connection_authentication_mode"`
	MetastoreDbConnectionPassword           string `tfschema:"metastore_db_connection_password"`
	MetastoreDbConnectionUrl                string `tfschema:"metastore_db_connection_url"`
	MetastoreDbConnectionUserName           string `tfschema:"metastore_db_connection_user_name"`
	MetastoreWarehouseDir                   string `tfschema:"metastore_warehouse_dir"`
}

type Coordinator struct {
	Debug                   []Debug `tfschema:"debug"`
	HighAvailabilityEnabled bool    `tfschema:"high_availability_enabled"`
}

type Debug struct {
	Enabled bool `tfschema:"enabled"`
	Port    int  `tfschema:"port"`
	Suspend bool `tfschema:"suspend"`
}

type Worker struct {
	Debug []Debug `tfschema:"debug"`
}

type UserPluginsSpec struct {
	Plugins []Plugin `tfschema:"plugins"`
}

type Plugin struct {
	Name    string `tfschema:"name"`
	Path    string `tfschema:"path"`
	Enabled bool   `tfschema:"enabled"`
}

type UserTelemetrySpec struct {
	HiveCatalogName          string `tfschema:"hive_catalog_name"`
	HiveCatalogSchema        string `tfschema:"hive_catalog_schema"`
	PartitionRetentionInDays int    `tfschema:"partition_retention_in_days"`
	Path                     string `tfschema:"path"`
}

func ClusterComputeProfileSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Required: true,
		MaxItems: 1,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"node": ClusterNodeSchema(),
			},
		},
	}
}

func ClusterNodeSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Required: true,
		MinItems: 1,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"count": {
					Type:     pluginsdk.TypeInt,
					Required: true,
				},

				"type": {
					Type:     pluginsdk.TypeString,
					Required: true,
				},

				"vm_size": {
					Type:     pluginsdk.TypeString,
					Required: true,
				},
			},
		},
	}
}

func ClusterProfileSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Required: true,
		MaxItems: 1,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"cluster_version": {
					Type:     pluginsdk.TypeString,
					Required: true,
				},

				"oss_version": {
					Type:     pluginsdk.TypeString,
					Required: true,
				},

				"authorization_profile": ClusterAuthorizationProfileSchema(),

				"autoscale_profile": ClusterAutoscaleProfileSchema(),

				"flink_profile": ClusterFlinkProfileSchema(),

				"identity_profile": ClusterIdentityProfileSchema(),

				"kafka_profile": ClusterKafkaProfileSchema(),

				"log_analytics_profile": ClusterLogAnalyticsProfileSchema(),

				"managed_identity_profile": ClusterManagedIdentityProfileSchema(),

				"cluster_access_profile": ClusterAccessProfileSchema(),

				"prometheus_profile": PrometheusProfileSchema(),

				"ranger_profile": RangerProfileSchema(),

				"script_action_profile": ScriptActionProfileSchema(),

				"secrets_profile": SecretsProfileSchema(),

				"service_configs_profiles": ServiceConfigsProfileSchema(),

				"spark_profile": SparkProfileSchema(),

				"ssh_profile": SshProfileSchema(),

				"trino_profile": TrinoProfileSchema(),
			},
		},
	}
}

func ClusterAuthorizationProfileSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Required: true,
		MaxItems: 1,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"group_ids": {
					Type:     pluginsdk.TypeList,
					Optional: true,
					Elem: &pluginsdk.Schema{
						Type: pluginsdk.TypeString,
					},
				},

				"user_ids": {
					Type:     pluginsdk.TypeList,
					Optional: true,
					Elem: &pluginsdk.Schema{
						Type: pluginsdk.TypeString,
					},
				},
			},
		},
	}
}

func ClusterLoadBasedConfigSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Optional: true,
		MaxItems: 1,
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
					Type:     pluginsdk.TypeInt,
					Optional: true,
				},

				"cooldown_period": {
					Type:     pluginsdk.TypeInt,
					Optional: true,
				},
			},
		},
	}
}

func ClusterScalingRuleSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Required: true,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"action_type": {
					Type:     pluginsdk.TypeString,
					Required: true,
					ValidateFunc: validation.StringInSlice([]string{
						string(hdinsights.ScaleActionTypeScaledown),
						string(hdinsights.ScaleActionTypeScaleup),
					}, false),
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
		Type:     pluginsdk.TypeList,
		Required: true,
		MaxItems: 1,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"operator": {
					Type:     pluginsdk.TypeString,
					Required: true,
					ValidateFunc: validation.StringInSlice([]string{
						string(hdinsights.ComparisonOperatorGreaterThan),
						string(hdinsights.ComparisonOperatorGreaterThanOrEqual),
						string(hdinsights.ComparisonOperatorLessThan),
						string(hdinsights.ComparisonOperatorLessThanOrEqual),
					}, false),
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
		Type:     pluginsdk.TypeList,
		Optional: true,
		MaxItems: 1,
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
		Type:     pluginsdk.TypeList,
		Required: true,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"count": {
					Type:     pluginsdk.TypeInt,
					Required: true,
				},

				"days": {
					Type:     pluginsdk.TypeList,
					Required: true,
					Elem: &pluginsdk.Schema{
						Type: pluginsdk.TypeString,
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

func ClusterAutoscaleProfileSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Optional: true,
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

func ClusterFlinkProfileSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Optional: true,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"job_manager": ClusterJobManagerSchema(),

				"task_manager": ClusterTaskManagerSchema(),

				"storage": ClusterStorageSchema(),

				"deployment_mode": {
					Type:     pluginsdk.TypeString,
					Optional: true,
					ValidateFunc: validation.StringInSlice([]string{
						string(hdinsights.DeploymentModeApplication),
						string(hdinsights.DeploymentModeSession),
					}, false),
				},

				"num_replicas": {
					Type:     pluginsdk.TypeInt,
					Required: true,
				},

				"catalog_option": ClusterCatalogOptionSchema(),

				"history_server": ClusterHistoryServerSchema(),

				"job_spec": ClusterJobSpecSchema(),
			},
		},
	}
}

func ClusterCatalogOptionSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"metastore_db_connection_uri": {
					Type:     pluginsdk.TypeString,
					Required: true,
				},

				"metastore_db_connection_authentication_mode": {
					Type:     pluginsdk.TypeString,
					Optional: true,
					ValidateFunc: validation.StringInSlice([]string{
						string(hdinsights.DbConnectionAuthenticationModeIdentityAuth),
						string(hdinsights.DbConnectionAuthenticationModeSqlAuth),
					}, false),
				},

				"metastore_db_connection_password": {
					Type:     pluginsdk.TypeString,
					Optional: true,
				},

				"metastore_db_connection_user_name": {
					Type:     pluginsdk.TypeString,
					Optional: true,
				},
			},
		},
	}
}

func ClusterHistoryServerSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"cpu": {
					Type:     pluginsdk.TypeFloat,
					Required: true,
				},

				"memory": {
					Type:     pluginsdk.TypeInt,
					Required: true,
				},
			},
		},
	}
}

func ClusterJobManagerSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Required: true,
		MaxItems: 1,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"cpu": {
					Type:     pluginsdk.TypeFloat,
					Required: true,
				},

				"memory": {
					Type:     pluginsdk.TypeInt,
					Required: true,
				},
			},
		},
	}
}

func ClusterJobSpecSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"jar_name": {
					Type:     pluginsdk.TypeString,
					Required: true,
				},

				"job_jar_directory": {
					Type:     pluginsdk.TypeString,
					Required: true,
				},

				"upgrade_mode": {
					Type:     pluginsdk.TypeString,
					Required: true,
					ValidateFunc: validation.StringInSlice([]string{
						string(hdinsights.UpgradeModeLASTSTATEUPDATE),
						string(hdinsights.UpgradeModeSTATELESSUPDATE),
						string(hdinsights.UpgradeModeUPDATE),
					}, false),
				},

				"save_point_name": {
					Type:     pluginsdk.TypeString,
					Optional: true,
				},

				"args": {
					Type:     pluginsdk.TypeString,
					Optional: true,
				},

				"entry_class": {
					Type:     pluginsdk.TypeString,
					Optional: true,
				},
			},
		},
	}
}

func ClusterStorageSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Required: true,
		MaxItems: 1,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"storage_url": {
					Type:     pluginsdk.TypeString,
					Required: true,
				},

				"storage_key": {
					Type:     pluginsdk.TypeString,
					Optional: true,
				},
			},
		},
	}
}

func ClusterTaskManagerSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Required: true,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"cpu": {
					Type:     pluginsdk.TypeFloat,
					Required: true,
				},

				"memory": {
					Type:     pluginsdk.TypeInt,
					Required: true,
				},
			},
		},
	}
}

func ClusterIdentityProfileSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"msi_client_id": {
					Type:     pluginsdk.TypeString,
					Required: true,
				},

				"msi_object_id": {
					Type:     pluginsdk.TypeString,
					Required: true,
				},

				"msi_resource_id": {
					Type:     pluginsdk.TypeString,
					Required: true,
				},
			},
		},
	}
}

func ClusterKafkaProfileSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"disk_storage": DiskStorageSchema(),

				"kraft_enabled": {
					Type:     pluginsdk.TypeBool,
					Optional: true,
				},

				"public_endpoint_enabled": {
					Type:     pluginsdk.TypeBool,
					Optional: true,
				},

				"remote_storage_uri": {
					Type:     pluginsdk.TypeString,
					Optional: true,
				},
			},
		},
	}
}

func DiskStorageSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Required: true,
		MaxItems: 1,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"disk_size": {
					Type:     pluginsdk.TypeInt,
					Required: true,
				},

				"disk_type": {
					Type:     pluginsdk.TypeString,
					Required: true,
					ValidateFunc: validation.StringInSlice([]string{
						string(hdinsights.DataDiskTypePremiumSSDLRS),
						string(hdinsights.DataDiskTypeStandardHDDLRS),
						string(hdinsights.DataDiskTypeStandardSSDLRS),
						string(hdinsights.DataDiskTypePremiumSSDVTwoLRS),
						string(hdinsights.DataDiskTypeStandardSSDZRS),
						string(hdinsights.DataDiskTypePremiumSSDZRS),
					}, false),
				},
			},
		},
	}
}

func ClusterLogAnalyticsProfileSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"enabled": {
					Type:     pluginsdk.TypeBool,
					Required: true,
				},

				"metric_enabled": {
					Type:     pluginsdk.TypeString,
					Optional: true,
				},

				"application_logs": {
					Type:     pluginsdk.TypeList,
					Optional: true,
					MaxItems: 1,
					Elem: &pluginsdk.Resource{
						Schema: map[string]*pluginsdk.Schema{
							"std_error_enabled": {
								Type:     pluginsdk.TypeBool,
								Optional: true,
							},

							"std_out_enabled": {
								Type:     pluginsdk.TypeBool,
								Optional: true,
							},
						},
					},
				},
			},
		},
	}
}

func ClusterManagedIdentityProfileSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"identities": ClusterIdentitySchema(),
			},
		},
	}
}

func ClusterIdentitySchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Required: true,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"client_id": {
					Type:     pluginsdk.TypeString,
					Required: true,
				},

				"object_id": {
					Type:     pluginsdk.TypeString,
					Required: true,
				},

				"resource_id": {
					Type:     pluginsdk.TypeString,
					Required: true,
				},

				"type": {
					Type:     pluginsdk.TypeString,
					Required: true,
					ValidateFunc: validation.StringInSlice([]string{
						string(hdinsights.ManagedIdentityTypeCluster),
						string(hdinsights.ManagedIdentityTypeInternal),
						string(hdinsights.ManagedIdentityTypeUser),
					}, false),
				},
			},
		},
	}
}

func ClusterAccessProfileSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"internal_ingress_enabled": {
					Type:     pluginsdk.TypeBool,
					Required: true,
				},
			},
		},
	}
}

func PrometheusProfileSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"enabled": {
					Type:     pluginsdk.TypeBool,
					Required: true,
				},
			},
		},
	}
}

func RangerProfileSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"ranger_admin": RangerAdminSchema(),

				"ranger_user_sync": RangerUserSyncSchema(),

				"storage_account": {
					Type:     pluginsdk.TypeString,
					Optional: true,
				},
			},
		},
	}
}

func ScriptActionProfileSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Optional: true,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"name": {
					Type:     pluginsdk.TypeString,
					Required: true,
				},

				"services": {
					Type:     pluginsdk.TypeList,
					Required: true,
					Elem: &pluginsdk.Schema{
						Type: pluginsdk.TypeString,
					},
				},

				"type": {
					Type:     pluginsdk.TypeString,
					Required: true,
				},

				"url": {
					Type:     pluginsdk.TypeString,
					Required: true,
				},

				"parameters": {
					Type:     pluginsdk.TypeString,
					Optional: true,
				},

				"should_persist": {
					Type:     pluginsdk.TypeBool,
					Required: true,
				},

				"timeout_in_minutes": {
					Type:     pluginsdk.TypeInt,
					Required: true,
				},
			},
		},
	}
}

func RangerAdminSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Required: true,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"admins": {
					Type:     pluginsdk.TypeList,
					Required: true,
					Elem: &pluginsdk.Schema{
						Type: pluginsdk.TypeString,
					},
				},

				"database": DatabaseSchema(),
			},
		},
	}
}

func DatabaseSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Required: true,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"host": {
					Type:     pluginsdk.TypeString,
					Required: true,
				},

				"name": {
					Type:     pluginsdk.TypeString,
					Required: true,
				},

				"password_secret_ref": {
					Type:     pluginsdk.TypeString,
					Optional: true,
				},

				"username": {
					Type:     pluginsdk.TypeString,
					Optional: true,
				},
			},
		},
	}
}

func RangerUserSyncSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Required: true,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"enabled": {
					Type:     pluginsdk.TypeBool,
					Optional: true,
				},

				"groups": {
					Type:     pluginsdk.TypeList,
					Optional: true,
					Elem: &pluginsdk.Schema{
						Type: pluginsdk.TypeString,
					},
				},

				"mode": {
					Type:     pluginsdk.TypeString,
					Optional: true,
					ValidateFunc: validation.StringInSlice([]string{
						string(hdinsights.RangerUsersyncModeStatic),
						string(hdinsights.RangerUsersyncModeAutomatic),
					}, false),
				},

				"user_mapping_location": {
					Type:     pluginsdk.TypeString,
					Optional: true,
				},

				"users": {
					Type:     pluginsdk.TypeList,
					Optional: true,
					Elem: &pluginsdk.Schema{
						Type: pluginsdk.TypeString,
					},
				},
			},
		},
	}
}

func SecretsProfileSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"key_vault_resource_id": {
					Type:     pluginsdk.TypeString,
					Required: true,
				},

				"secrets": SecretSchema(),
			},
		},
	}
}

func SecretSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Optional: true,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"key_vault_object_name": {
					Type:     pluginsdk.TypeString,
					Required: true,
				},

				"type": {
					Type:     pluginsdk.TypeString,
					Required: true,
					ValidateFunc: validation.StringInSlice([]string{
						string(hdinsights.KeyVaultObjectTypeCertificate),
						string(hdinsights.KeyVaultObjectTypeKey),
						string(hdinsights.KeyVaultObjectTypeSecret),
					}, false),
				},

				"reference_name": {
					Type:     pluginsdk.TypeString,
					Required: true,
				},

				"version": {
					Type:     pluginsdk.TypeString,
					Optional: true,
				},
			},
		},
	}
}

func ServiceConfigsProfileSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Optional: true,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"service_name": {
					Type:     pluginsdk.TypeString,
					Required: true,
				},

				"configs": ConfigSchema(),
			},
		},
	}
}

func ConfigSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Required: true,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"component": {
					Type:     pluginsdk.TypeString,
					Required: true,
				},

				"files": FileSchema(),
			},
		},
	}
}

func FileSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Required: true,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"file_name": {
					Type:     pluginsdk.TypeString,
					Required: true,
				},

				"content": {
					Type:     pluginsdk.TypeString,
					Optional: true,
				},

				"encoding": {
					Type:     pluginsdk.TypeString,
					Optional: true,
					ValidateFunc: validation.StringInSlice([]string{
						string(hdinsights.ContentEncodingBaseSixFour),
						string(hdinsights.ContentEncodingNone),
					}, false),
				},

				"path": {
					Type:     pluginsdk.TypeString,
					Optional: true,
				},

				"values": {
					Type:     pluginsdk.TypeMap,
					Optional: true,
				},
			},
		},
	}
}

func SparkProfileSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Optional: true,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"default_storage_url": {
					Type:     pluginsdk.TypeString,
					Optional: true,
				},

				"metastore_spec": MetastoreSpecSchema(),
			},
		},
	}
}

func MetastoreSpecSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Optional: true,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"db_name": {
					Type:     pluginsdk.TypeString,
					Required: true,
				},

				"db_server_host": {
					Type:     pluginsdk.TypeString,
					Required: true,
				},

				"db_connection_authentication_mode": {
					Type:     pluginsdk.TypeString,
					Optional: true,
				},

				"db_password_secret_name": {
					Type:     pluginsdk.TypeString,
					Optional: true,
				},

				"db_user_name": {
					Type:     pluginsdk.TypeString,
					Optional: true,
				},

				"key_vault_id": {
					Type:     pluginsdk.TypeString,
					Optional: true,
				},

				"thrift_url": {
					Type:     pluginsdk.TypeString,
					Optional: true,
				},
			},
		},
	}
}

func SshProfileSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Optional: true,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"count": {
					Type:     pluginsdk.TypeInt,
					Required: true,
				},

				"vm_size": {
					Type:     pluginsdk.TypeString,
					Optional: true,
				},
			},
		},
	}
}

func TrinoProfileSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Optional: true,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"catalog_options": {
					Type:     pluginsdk.TypeList,
					Optional: true,
					Elem: &pluginsdk.Resource{
						Schema: map[string]*pluginsdk.Schema{
							"hive": TrinoCatalogOptionHiveSchema(),
						},
					},
				},

				"coordinator": CoordinatorSchema(),

				"worker": WorkerSchema(),

				"user_plugins_spec": UserPluginsSpecSchema(),

				"user_telemetry_spec": UserTelemetrySpecSchema(),
			},
		},
	}
}

func TrinoCatalogOptionHiveSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Optional: true,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"catalog_name": {
					Type:     pluginsdk.TypeString,
					Required: true,
				},

				"metastore_db_connection_url": {
					Type:     pluginsdk.TypeString,
					Required: true,
				},

				"metastore_warehouse_dir": {
					Type:     pluginsdk.TypeString,
					Required: true,
				},

				"metastore_db_connection_user_name": {
					Type:     pluginsdk.TypeString,
					Optional: true,
				},

				"metastore_db_connection_authentication_mode": {
					Type:     pluginsdk.TypeString,
					Optional: true,
				},

				"metastore_db_connection_password": {
					Type:     pluginsdk.TypeString,
					Optional: true,
				},
			},
		},
	}
}

func CoordinatorSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"debug": DebugSchema(),

				"high_availability_enabled": {
					Type:     pluginsdk.TypeBool,
					Optional: true,
				},
			},
		},
	}
}

func DebugSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"enabled": {
					Type:     pluginsdk.TypeBool,
					Optional: true,
				},

				"port": {
					Type:     pluginsdk.TypeInt,
					Optional: true,
				},

				"suspend": {
					Type:     pluginsdk.TypeBool,
					Optional: true,
				},
			},
		},
	}
}

func WorkerSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"debug": DebugSchema(),
			},
		},
	}
}

func UserPluginsSpecSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"plugins": PluginSchema(),
			},
		},
	}
}

func PluginSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Optional: true,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"enabled": {
					Type:     pluginsdk.TypeBool,
					Optional: true,
				},

				"name": {
					Type:     pluginsdk.TypeString,
					Optional: true,
				},

				"path": {
					Type:     pluginsdk.TypeString,
					Optional: true,
				},
			},
		},
	}
}

func UserTelemetrySpecSchema() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Type:     pluginsdk.TypeList,
		Optional: true,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"hive_catalog_name": {
					Type:     pluginsdk.TypeString,
					Optional: true,
				},

				"hive_catalog_schema": {
					Type:     pluginsdk.TypeString,
					Optional: true,
				},

				"partition_retention_in_days": {
					Type:     pluginsdk.TypeInt,
					Optional: true,
				},

				"path": {
					Type:     pluginsdk.TypeString,
					Optional: true,
				},
			},
		},
	}
}

func expandClusterProfile(input []ClusterProfile) hdinsights.ClusterProfile {
	if len(input) == 0 {
		return hdinsights.ClusterProfile{}
	}

	profile := input[0]

	return hdinsights.ClusterProfile{
		AuthorizationProfile:   expandClusterAuthorizationProfile(profile.AuthorizationProfile),
		AutoscaleProfile:       expandClusterAutoscaleProfile(profile.AutoscaleProfile),
		FlinkProfile:           expandClusterFlinkProfile(profile.FlinkProfile),
		IdentityProfile:        expandClusterIdentityProfile(profile.IdentityProfile),
		KafkaProfile:           expandClusterKafkaProfile(profile.KafkaProfile),
		LogAnalyticsProfile:    expandClusterLogAnalyticsProfile(profile.LogAnalyticsProfile),
		ManagedIdentityProfile: expandClusterManagedIdentityProfile(profile.ManagedIdentityProfile),
		PrometheusProfile:      expandClusterPrometheusProfile(profile.PrometheusProfile),
		RangerProfile:          expandRangerProfile(profile.RangerProfile),
		ScriptActionProfiles:   expandScriptActionProfile(profile.ScriptActionProfile),
		SecretsProfile:         expandSecretsProfile(profile.SecretsProfile),
		ServiceConfigsProfiles: expandServiceConfigsProfiles(profile.ServiceConfigsProfiles),
		SshProfile:             expandSshProfile(profile.SshProfile),
		TrinoProfile:           expandTrinoProfile(profile.TrinoProfile),
	}
}

func expandClusterProfilePatch(input []ClusterProfile) *hdinsights.UpdatableClusterProfile {
	if len(input) == 0 {
		return nil
	}

	profile := input[0]

	return &hdinsights.UpdatableClusterProfile{
		AuthorizationProfile:   pointer.To(expandClusterAuthorizationProfile(profile.AuthorizationProfile)),
		AutoscaleProfile:       expandClusterAutoscaleProfile(profile.AutoscaleProfile),
		LogAnalyticsProfile:    expandClusterLogAnalyticsProfile(profile.LogAnalyticsProfile),
		PrometheusProfile:      expandPrometheusProfile(profile.PrometheusProfile),
		RangerProfile:          expandRangerProfile(profile.RangerProfile),
		ScriptActionProfiles:   expandScriptActionProfile(profile.ScriptActionProfile),
		SecretsProfile:         expandSecretsProfile(profile.SecretsProfile),
		ServiceConfigsProfiles: expandServiceConfigsProfiles(profile.ServiceConfigsProfiles),
		SshProfile:             expandSshProfile(profile.SshProfile),
		TrinoProfile:           expandTrinoProfile(profile.TrinoProfile),
	}
}

func expandClusterPrometheusProfile(input []PrometheusProfile) *hdinsights.ClusterPrometheusProfile {
	if len(input) == 0 {
		return nil
	}

	profile := input[0]

	return &hdinsights.ClusterPrometheusProfile{
		Enabled: profile.Enabled,
	}
}

func expandSshProfile(input []SshProfile) *hdinsights.SshProfile {
	if len(input) == 0 {
		return nil
	}

	profile := input[0]

	return &hdinsights.SshProfile{
		Count:  int64(profile.Count),
		VMSize: pointer.To(profile.VmSize),
	}
}

func expandSecretsProfile(input []SecretsProfile) *hdinsights.SecretsProfile {
	if len(input) == 0 {
		return nil
	}

	profile := input[0]

	return &hdinsights.SecretsProfile{
		KeyVaultResourceId: profile.KeyVaultResourceId,
		Secrets:            expandSecrets(profile.Secrets),
	}
}

func expandSecrets(input []Secret) *[]hdinsights.SecretReference {
	secrets := make([]hdinsights.SecretReference, 0)

	for _, secret := range input {
		secrets = append(secrets, hdinsights.SecretReference{
			KeyVaultObjectName: secret.KeyVaultObjectName,
			ReferenceName:      secret.ReferenceName,
			Type:               hdinsights.KeyVaultObjectType(secret.Type),
			Version:            pointer.To(secret.Version),
		})
	}

	return pointer.To(secrets)
}

func expandScriptActionProfile(input []ScriptActionProfile) *[]hdinsights.ScriptActionProfile {
	actions := make([]hdinsights.ScriptActionProfile, 0)

	for _, action := range input {
		actions = append(actions, hdinsights.ScriptActionProfile{
			Name:             action.Name,
			Services:         action.Services,
			Type:             action.Type,
			Url:              action.Url,
			Parameters:       pointer.To(action.Parameters),
			ShouldPersist:    pointer.To(action.ShouldPersist),
			TimeoutInMinutes: pointer.To(int64(action.TimeoutInMinutes)),
		})
	}

	return pointer.To(actions)
}

func expandServiceConfigsProfiles(input []ServiceConfigsProfile) *[]hdinsights.ClusterServiceConfigsProfile {
	configs := make([]hdinsights.ClusterServiceConfigsProfile, 0)

	for _, config := range input {
		configs = append(configs, hdinsights.ClusterServiceConfigsProfile{
			ServiceName: config.ServiceName,
			Configs:     expandConfigs(config.Configs),
		})
	}

	return pointer.To(configs)
}

func expandConfigs(input []Config) []hdinsights.ClusterServiceConfig {
	configs := make([]hdinsights.ClusterServiceConfig, 0)

	for _, config := range input {
		configs = append(configs, hdinsights.ClusterServiceConfig{
			Component: config.Component,
			Files:     expandFiles(config.Files),
		})
	}

	return configs
}

func expandFiles(input []File) []hdinsights.ClusterConfigFile {
	files := make([]hdinsights.ClusterConfigFile, 0)

	for _, file := range input {
		files = append(files, hdinsights.ClusterConfigFile{
			FileName: file.FileName,
			Content:  pointer.To(file.Content),
			Encoding: pointer.To(hdinsights.ContentEncoding(file.Encoding)),
			Path:     pointer.To(file.Path),
			Values:   pointer.To(file.Values),
		})
	}

	return files
}

func expandTrinoProfile(input []TrinoProfile) *hdinsights.TrinoProfile {
	if len(input) == 0 {
		return nil
	}

	profile := input[0]

	return &hdinsights.TrinoProfile{
		CatalogOptions:    expandTrinoCatalogOption(profile.CatalogOptions),
		Coordinator:       expandTrinoCoordinator(profile.Coordinator),
		Worker:            expandTrinoWorker(profile.Worker),
		UserPluginsSpec:   expandTrinoUserPluginsSpec(profile.UserPluginsSpec),
		UserTelemetrySpec: expandTrinoUserTelemetrySpec(profile.UserTelemetrySpec),
	}
}

func expandTrinoUserPluginsSpec(input []UserPluginsSpec) *hdinsights.TrinoUserPlugins {
	if len(input) == 0 {
		return nil
	}

	spec := input[0]

	return &hdinsights.TrinoUserPlugins{
		Plugins: expandTrinoPlugins(spec.Plugins),
	}
}

func expandTrinoPlugins(input []Plugin) *[]hdinsights.TrinoUserPlugin {
	if len(input) == 0 {
		return nil
	}

	plugins := make([]hdinsights.TrinoUserPlugin, 0)

	for _, plugin := range input {
		plugins = append(plugins, hdinsights.TrinoUserPlugin{
			Enabled: pointer.To(plugin.Enabled),
			Name:    pointer.To(plugin.Name),
			Path:    pointer.To(plugin.Path),
		})
	}

	return pointer.To(plugins)
}

func expandTrinoUserTelemetrySpec(input []UserTelemetrySpec) *hdinsights.TrinoUserTelemetry {
	if len(input) == 0 {
		return nil
	}

	spec := input[0]

	return &hdinsights.TrinoUserTelemetry{
		Storage: &hdinsights.TrinoTelemetryConfig{
			HivecatalogName:          pointer.To(spec.HiveCatalogName),
			HivecatalogSchema:        pointer.To(spec.HiveCatalogSchema),
			PartitionRetentionInDays: pointer.To(int64(spec.PartitionRetentionInDays)),
			Path:                     pointer.To(spec.Path),
		},
	}
}

func expandTrinoWorker(input []Worker) *hdinsights.TrinoWorker {
	if len(input) == 0 {
		return nil
	}

	worker := input[0]

	return &hdinsights.TrinoWorker{
		Debug: expandDebug(worker.Debug),
	}
}

func expandTrinoCatalogOption(input []TrinoCatalogOption) *hdinsights.CatalogOptions {
	if len(input) == 0 {
		return nil
	}

	option := input[0]

	return &hdinsights.CatalogOptions{
		Hive: expandTrinoCatalogOptionHive(option.Hive),
	}
}

func expandTrinoCatalogOptionHive(input []Hive) *[]hdinsights.HiveCatalogOption {
	if len(input) == 0 {
		return nil
	}

	result := make([]hdinsights.HiveCatalogOption, 0)
	for _, hive := range input {
		result = append(result, hdinsights.HiveCatalogOption{
			CatalogName:                             hive.CatalogName,
			MetastoreDbConnectionURL:                hive.MetastoreDbConnectionUrl,
			MetastoreWarehouseDir:                   hive.MetastoreWarehouseDir,
			MetastoreDbConnectionUserName:           pointer.To(hive.MetastoreDbConnectionUserName),
			MetastoreDbConnectionAuthenticationMode: pointer.To(hdinsights.MetastoreDbConnectionAuthenticationMode(hive.MetastoreDbConnectionAuthenticationMode)),
			MetastoreDbConnectionPasswordSecret:     pointer.To(hive.MetastoreDbConnectionPassword),
		})
	}

	return pointer.To(result)
}

func expandTrinoCoordinator(input []Coordinator) *hdinsights.TrinoCoordinator {
	if len(input) == 0 {
		return nil
	}

	coordinator := input[0]

	return &hdinsights.TrinoCoordinator{
		Debug:                   expandDebug(coordinator.Debug),
		HighAvailabilityEnabled: pointer.To(coordinator.HighAvailabilityEnabled),
	}
}

func expandDebug(input []Debug) *hdinsights.TrinoDebugConfig {
	if len(input) == 0 {
		return nil
	}

	debug := input[0]

	return &hdinsights.TrinoDebugConfig{
		Enable:  pointer.To(debug.Enabled),
		Port:    pointer.To(int64(debug.Port)),
		Suspend: pointer.To(debug.Suspend),
	}
}

func expandPrometheusProfile(input []PrometheusProfile) *hdinsights.ClusterPrometheusProfile {
	if len(input) == 0 {
		return nil
	}
	profile := input[0]

	return &hdinsights.ClusterPrometheusProfile{
		Enabled: profile.Enabled,
	}
}

func expandRangerProfile(input []RangerProfile) *hdinsights.RangerProfile {
	if len(input) == 0 {
		return nil
	}
	profile := input[0]

	result := &hdinsights.RangerProfile{
		RangerAdmin:    expandRangerAdmin(profile.RangerAdmin),
		RangerUsersync: expandRangerUserSync(profile.RangerUserSync),
	}
	if profile.StorageAccount != "" {
		result.RangerAudit = pointer.To(hdinsights.RangerAuditSpec{
			StorageAccount: pointer.To(profile.StorageAccount),
		})
	}

	return result
}

func expandRangerAdmin(input []RangerAdmin) hdinsights.RangerAdminSpec {
	admin := input[0]

	return hdinsights.RangerAdminSpec{
		Admins:   admin.Admins,
		Database: expandDatabase(admin.Database),
	}
}

func expandDatabase(input []Database) hdinsights.RangerAdminSpecDatabase {
	if len(input) == 0 {
		return hdinsights.RangerAdminSpecDatabase{}
	}
	database := input[0]

	return hdinsights.RangerAdminSpecDatabase{
		Host:              database.Host,
		Name:              database.Name,
		PasswordSecretRef: pointer.To(database.PasswordSecretRef),
		Username:          pointer.To(database.Username),
	}
}

func expandRangerUserSync(input []RangerUserSync) hdinsights.RangerUsersyncSpec {
	userSync := input[0]

	return hdinsights.RangerUsersyncSpec{
		Enabled:             pointer.To(userSync.Enabled),
		Groups:              pointer.To(userSync.Groups),
		Mode:                pointer.To(hdinsights.RangerUsersyncMode(userSync.Mode)),
		UserMappingLocation: pointer.To(userSync.UserMappingLocation),
		Users:               pointer.To(userSync.Users),
	}
}

func expandClusterAuthorizationProfile(input []AuthorizationProfile) hdinsights.AuthorizationProfile {
	if len(input) == 0 {
		return hdinsights.AuthorizationProfile{}
	}
	profile := input[0]

	return hdinsights.AuthorizationProfile{
		GroupIds: pointer.To(profile.GroupIds),
		UserIds:  pointer.To(profile.UserIds),
	}
}

func expandClusterAutoscaleProfile(input []AutoscaleProfile) *hdinsights.AutoscaleProfile {
	if len(input) == 0 {
		return nil
	}
	profile := input[0]

	return &hdinsights.AutoscaleProfile{
		AutoscaleType:               pointer.To(hdinsights.AutoscaleType(profile.AutoscaleType)),
		Enabled:                     profile.AutoscaleEnabled,
		GracefulDecommissionTimeout: pointer.To(int64(profile.GracefulDecommissionTimeout)),
		LoadBasedConfig:             expandClusterLoadBasedConfig(profile.LoadBasedConfig),
		ScheduleBasedConfig:         expandClusterScheduleBasedConfig(profile.ScheduleBasedConfig),
	}
}

func expandClusterFlinkProfile(input []FlinkProfile) *hdinsights.FlinkProfile {
	if len(input) == 0 {
		return nil
	}
	profile := input[0]

	return &hdinsights.FlinkProfile{
		DeploymentMode: pointer.To(hdinsights.DeploymentMode(profile.DeploymentMode)),
		NumReplicas:    pointer.To(int64(profile.NumReplicas)),
		CatalogOptions: expandClusterCatalogOption(profile.CatalogOptions),
		HistoryServer:  expandComputeResource(profile.HistoryServer),
		JobManager:     pointer.From(expandComputeResource(profile.JobManager)),
		JobSpec:        expandClusterJobSpec(profile.JobSpec),
		Storage:        expandClusterStorage(profile.Storage),
		TaskManager:    pointer.From(expandComputeResource(profile.TaskManager)),
	}
}
func expandClusterLoadBasedConfig(input []LoadBasedConfig) *hdinsights.LoadBasedConfig {
	if len(input) == 0 {
		return nil
	}
	config := input[0]

	return &hdinsights.LoadBasedConfig{
		CooldownPeriod: pointer.To(int64(config.CooldownPeriod)),
		MaxNodes:       int64(config.MaxNodes),
		MinNodes:       int64(config.MinNodes),
		PollInterval:   pointer.To(int64(config.PollInterval)),
		ScalingRules:   expandClusterScalingRule(config.ScalingRules),
	}
}

func expandClusterScalingRule(input []ScalingRule) []hdinsights.ScalingRule {
	rules := make([]hdinsights.ScalingRule, 0)

	for _, rule := range input {
		rules = append(rules, hdinsights.ScalingRule{
			ActionType:      hdinsights.ScaleActionType(rule.ActionType),
			EvaluationCount: int64(rule.EvaluationCount),
			ScalingMetric:   rule.ScalingMetric,
			ComparisonRule:  expandClusterComparisonRule(rule.ComparisonRule),
		})
	}

	return rules
}

func expandClusterComparisonRule(input []ComparisonRule) hdinsights.ComparisonRule {
	if len(input) == 0 {
		return hdinsights.ComparisonRule{}
	}
	rule := input[0]

	return hdinsights.ComparisonRule{
		Operator:  hdinsights.ComparisonOperator(rule.Operator),
		Threshold: rule.Threshold,
	}
}

func expandClusterScheduleBasedConfig(input []ScheduleBasedConfig) *hdinsights.ScheduleBasedConfig {
	if len(input) == 0 {
		return nil
	}
	config := input[0]

	return &hdinsights.ScheduleBasedConfig{
		DefaultCount: int64(config.DefaultCount),
		Schedules:    expandClusterSchedule(config.Schedules),
		TimeZone:     config.TimeZone,
	}
}

func expandClusterSchedule(input []Schedule) []hdinsights.Schedule {
	schedules := make([]hdinsights.Schedule, 0)

	for _, value := range input {
		schedule := hdinsights.Schedule{
			Count:     int64(value.Count),
			EndTime:   value.EndTime,
			StartTime: value.StartTime,
		}
		for _, day := range value.Days {
			schedule.Days = append(schedule.Days, hdinsights.ScheduleDay(day))
		}
		schedules = append(schedules, schedule)
	}

	return schedules
}

func expandClusterCatalogOption(input []FlinkCatalogOption) *hdinsights.FlinkCatalogOptions {
	if len(input) == 0 {
		return nil
	}
	option := input[0]
	hive := hdinsights.FlinkHiveCatalogOption{
		MetastoreDbConnectionAuthenticationMode: pointer.To(hdinsights.MetastoreDbConnectionAuthenticationMode(option.MetastoreDbConnectionAuthenticationMode)),
		MetastoreDbConnectionPasswordSecret:     pointer.To(option.MetastoreDbConnectionPassword),
		MetastoreDbConnectionURL:                option.MetastoreDbConnectionUrl,
		MetastoreDbConnectionUserName:           pointer.To(option.MetastoreDbConnectionUserName),
	}
	return &hdinsights.FlinkCatalogOptions{
		Hive: pointer.To(hive),
	}
}

func expandComputeResource(input []ComputeResource) *hdinsights.ComputeResourceDefinition {
	if len(input) == 0 {
		return nil
	}
	server := input[0]

	return &hdinsights.ComputeResourceDefinition{
		Cpu:    server.Cpu,
		Memory: int64(server.Memory),
	}
}

func expandClusterJobSpec(input []JobSpec) *hdinsights.FlinkJobProfile {
	if len(input) == 0 {
		return nil
	}
	spec := input[0]

	return &hdinsights.FlinkJobProfile{
		Args:            pointer.To(spec.Args),
		EntryClass:      pointer.To(spec.EntryClass),
		JarName:         spec.JarName,
		JobJarDirectory: spec.JobJarDirectory,
		UpgradeMode:     hdinsights.UpgradeMode(spec.UpgradeMode),
		SavePointName:   pointer.To(spec.SavePointName),
	}
}

func expandClusterIdentityProfile(input []IdentityProfile) *hdinsights.IdentityProfile {
	if len(input) == 0 {
		return nil
	}
	profile := input[0]

	return &hdinsights.IdentityProfile{
		MsiClientId:   profile.MsiClientId,
		MsiObjectId:   profile.MsiObjectId,
		MsiResourceId: profile.MsiResourceId,
	}
}

func expandClusterKafkaProfile(input []KafkaProfile) *hdinsights.KafkaProfile {
	if len(input) == 0 {
		return nil
	}
	profile := input[0]

	return &hdinsights.KafkaProfile{
		DiskStorage:           expandDiskStorage(profile.DiskStorage),
		EnableKRaft:           pointer.To(profile.KRaftEnabled),
		EnablePublicEndpoints: pointer.To(profile.PublicEndpointEnabled),
		RemoteStorageUri:      pointer.To(profile.RemoteStorageUri),
	}
}

func expandDiskStorage(input []DiskStorage) hdinsights.DiskStorageProfile {
	if len(input) == 0 {
		return hdinsights.DiskStorageProfile{}
	}
	storage := input[0]

	return hdinsights.DiskStorageProfile{
		DataDiskSize: int64(storage.DiskSize),
		DataDiskType: hdinsights.DataDiskType(storage.DiskType),
	}
}

func expandClusterLogAnalyticsProfile(input []ClusterLogAnalyticsProfile) *hdinsights.ClusterLogAnalyticsProfile {
	if len(input) == 0 {
		return nil
	}
	profile := input[0]

	return &hdinsights.ClusterLogAnalyticsProfile{
		ApplicationLogs: expandApplicationLogs(profile.ApplicationLogs),
		Enabled:         profile.Enabled,
		MetricsEnabled:  pointer.To(profile.MetricsEnabled),
	}
}

func expandApplicationLogs(input []ApplicationLogs) *hdinsights.ClusterLogAnalyticsApplicationLogs {
	if len(input) == 0 {
		return nil
	}
	logs := input[0]

	return &hdinsights.ClusterLogAnalyticsApplicationLogs{
		StdErrorEnabled: pointer.To(logs.StdErrorEnabled),
		StdOutEnabled:   pointer.To(logs.StdOutEnabled),
	}
}

func expandClusterManagedIdentityProfile(input []ManagedIdentityProfile) *hdinsights.ManagedIdentityProfile {
	if len(input) == 0 {
		return nil
	}
	profile := input[0]

	return &hdinsights.ManagedIdentityProfile{
		IdentityList: expandIdentity(profile.Identities),
	}
}

func expandIdentity(input []Identity) []hdinsights.ManagedIdentitySpec {
	identities := make([]hdinsights.ManagedIdentitySpec, 0)

	for _, value := range input {
		identities = append(identities, hdinsights.ManagedIdentitySpec{
			ClientId:   value.ClientId,
			ObjectId:   value.ObjectId,
			ResourceId: value.ResourceId,
			Type:       hdinsights.ManagedIdentityType(value.Type),
		})
	}

	return identities
}

func expandClusterStorage(input []Storage) hdinsights.FlinkStorageProfile {
	if len(input) == 0 {
		return hdinsights.FlinkStorageProfile{}
	}
	storage := input[0]

	return hdinsights.FlinkStorageProfile{
		StorageUri: storage.StorageUrl,
		Storagekey: pointer.To(storage.StorageKey),
	}
}

func expandClusterComputeProfile(input []ClusterComputeProfile) *hdinsights.ComputeProfile {
	if len(input) == 0 {
		return nil
	}
	profile := input[0]

	return &hdinsights.ComputeProfile{
		Nodes: expandClusterNodes(profile.Node),
	}
}

func expandClusterNodes(input []Node) []hdinsights.NodeProfile {
	nodes := make([]hdinsights.NodeProfile, 0)

	for _, value := range input {
		nodes = append(nodes, hdinsights.NodeProfile{
			Count:  int64(value.Count),
			VMSize: value.VmSize,
			Type:   value.Type,
		})
	}

	return nodes
}

func flattenClusterProfile(input hdinsights.ClusterProfile) []ClusterProfile {
	return []ClusterProfile{
		{
			AuthorizationProfile:   flattenClusterAuthorizationProfile(input.AuthorizationProfile),
			AutoscaleProfile:       flattenClusterAutoscaleProfile(input.AutoscaleProfile),
			FlinkProfile:           flattenClusterFlinkProfile(input.FlinkProfile),
			IdentityProfile:        flattenClusterIdentityProfile(input.IdentityProfile),
			KafkaProfile:           flattenClusterKafkaProfile(input.KafkaProfile),
			LogAnalyticsProfile:    flattenClusterLogAnalyticsProfile(input.LogAnalyticsProfile),
			ManagedIdentityProfile: flattenClusterManagedIdentityProfile(input.ManagedIdentityProfile),
		},
	}
}

func flattenClusterAuthorizationProfile(input hdinsights.AuthorizationProfile) []AuthorizationProfile {
	result := make([]AuthorizationProfile, 0)

	profile := AuthorizationProfile{
		GroupIds: pointer.From(input.GroupIds),
		UserIds:  pointer.From(input.UserIds),
	}
	result = append(result, profile)

	return result
}

func flattenClusterAutoscaleProfile(input *hdinsights.AutoscaleProfile) []AutoscaleProfile {
	if input == nil {
		return []AutoscaleProfile{}
	}

	result := make([]AutoscaleProfile, 0)
	profile := AutoscaleProfile{
		AutoscaleEnabled: input.Enabled,
	}
	if input.AutoscaleType != nil {
		profile.AutoscaleType = string(pointer.From(input.AutoscaleType))
	}
	result = append(result, profile)

	return result
}

func flattenClusterFlinkProfile(input *hdinsights.FlinkProfile) []FlinkProfile {
	if input == nil {
		return []FlinkProfile{}
	}

	result := make([]FlinkProfile, 0)
	profile := FlinkProfile{
		JobManager:  flattenComputeResource(input.JobManager),
		TaskManager: flattenComputeResource(input.TaskManager),
		Storage:     flattenClusterStorage(input.Storage),
	}
	if input.DeploymentMode != nil {
		profile.DeploymentMode = string(pointer.From(input.DeploymentMode))
	}
	if input.NumReplicas != nil {
		profile.NumReplicas = int(pointer.From(input.NumReplicas))
	}
	if input.CatalogOptions != nil {
		profile.CatalogOptions = flattenClusterCatalogOption(input.CatalogOptions)
	}
	if input.HistoryServer != nil {
		profile.HistoryServer = flattenComputeResource(pointer.From(input.HistoryServer))
	}
	if input.JobSpec != nil {
		profile.JobSpec = flattenClusterJobSpec(input.JobSpec)
	}
	result = append(result, profile)

	return result
}

func flattenClusterCatalogOption(input *hdinsights.FlinkCatalogOptions) []FlinkCatalogOption {
	if input == nil {
		return []FlinkCatalogOption{}
	}

	result := make([]FlinkCatalogOption, 0)
	option := FlinkCatalogOption{
		MetastoreDbConnectionUrl: input.Hive.MetastoreDbConnectionURL,
	}
	if input.Hive.MetastoreDbConnectionAuthenticationMode != nil {
		option.MetastoreDbConnectionAuthenticationMode = string(pointer.From(input.Hive.MetastoreDbConnectionAuthenticationMode))
	}
	if input.Hive.MetastoreDbConnectionPasswordSecret != nil {
		option.MetastoreDbConnectionPassword = pointer.From(input.Hive.MetastoreDbConnectionPasswordSecret)
	}
	if input.Hive.MetastoreDbConnectionUserName != nil {
		option.MetastoreDbConnectionUserName = pointer.From(input.Hive.MetastoreDbConnectionUserName)
	}
	result = append(result, option)

	return result
}

func flattenComputeResource(input hdinsights.ComputeResourceDefinition) []ComputeResource {
	result := make([]ComputeResource, 0)
	server := ComputeResource{
		Cpu:    input.Cpu,
		Memory: int(input.Memory),
	}
	result = append(result, server)

	return result
}

func flattenClusterStorage(input hdinsights.FlinkStorageProfile) []Storage {
	result := make([]Storage, 0)
	storage := Storage{
		StorageUrl: input.StorageUri,
	}

	if input.Storagekey != nil {
		storage.StorageKey = pointer.From(input.Storagekey)
	}
	result = append(result, storage)

	return result
}

func flattenClusterJobSpec(input *hdinsights.FlinkJobProfile) []JobSpec {
	if input == nil {
		return []JobSpec{}
	}
	result := make([]JobSpec, 0)
	spec := JobSpec{
		Args:            pointer.From(input.Args),
		EntryClass:      pointer.From(input.EntryClass),
		JarName:         input.JarName,
		JobJarDirectory: input.JobJarDirectory,
		UpgradeMode:     string(input.UpgradeMode),
	}
	if input.SavePointName != nil {
		spec.SavePointName = pointer.From(input.SavePointName)
	}
	result = append(result, spec)

	return result
}

func flattenClusterIdentityProfile(input *hdinsights.IdentityProfile) []IdentityProfile {
	if input == nil {
		return []IdentityProfile{}
	}

	result := make([]IdentityProfile, 0)
	profile := IdentityProfile{
		MsiClientId:   input.MsiClientId,
		MsiObjectId:   input.MsiObjectId,
		MsiResourceId: input.MsiResourceId,
	}
	result = append(result, profile)

	return result
}

func flattenClusterKafkaProfile(input *hdinsights.KafkaProfile) []KafkaProfile {
	if input == nil {
		return []KafkaProfile{}
	}

	result := make([]KafkaProfile, 0)
	profile := KafkaProfile{
		DiskStorage: flattenDiskStorage(input.DiskStorage),
	}
	if input.EnableKRaft != nil {
		profile.KRaftEnabled = pointer.From(input.EnableKRaft)
	}
	if input.EnablePublicEndpoints != nil {
		profile.PublicEndpointEnabled = pointer.From(input.EnablePublicEndpoints)
	}
	if input.RemoteStorageUri != nil {
		profile.RemoteStorageUri = pointer.From(input.RemoteStorageUri)
	}

	result = append(result, profile)

	return result
}

func flattenDiskStorage(input hdinsights.DiskStorageProfile) []DiskStorage {
	result := make([]DiskStorage, 0)
	storage := DiskStorage{
		DiskSize: int(input.DataDiskSize),
		DiskType: string(input.DataDiskType),
	}
	result = append(result, storage)

	return result
}

func flattenClusterLogAnalyticsProfile(input *hdinsights.ClusterLogAnalyticsProfile) []ClusterLogAnalyticsProfile {
	if input == nil {
		return []ClusterLogAnalyticsProfile{}
	}

	result := make([]ClusterLogAnalyticsProfile, 0)
	profile := ClusterLogAnalyticsProfile{
		Enabled: input.Enabled,
	}
	if input.ApplicationLogs != nil {
		profile.ApplicationLogs = flattenApplicationLogs(input.ApplicationLogs)
	}
	if input.MetricsEnabled != nil {
		profile.MetricsEnabled = pointer.From(input.MetricsEnabled)
	}
	result = append(result, profile)

	return result
}

func flattenApplicationLogs(input *hdinsights.ClusterLogAnalyticsApplicationLogs) []ApplicationLogs {
	if input == nil {
		return []ApplicationLogs{}
	}

	result := make([]ApplicationLogs, 0)
	logs := ApplicationLogs{
		StdErrorEnabled: pointer.From(input.StdErrorEnabled),
		StdOutEnabled:   pointer.From(input.StdOutEnabled),
	}
	result = append(result, logs)

	return result
}

func flattenClusterManagedIdentityProfile(input *hdinsights.ManagedIdentityProfile) []ManagedIdentityProfile {
	if input == nil {
		return []ManagedIdentityProfile{}
	}

	result := make([]ManagedIdentityProfile, 0)
	profile := ManagedIdentityProfile{
		Identities: flattenIdentity(input.IdentityList),
	}
	result = append(result, profile)

	return result
}

func flattenIdentity(input []hdinsights.ManagedIdentitySpec) []Identity {
	identities := make([]Identity, 0)

	for _, value := range input {
		identities = append(identities, Identity{
			ClientId:   value.ClientId,
			ObjectId:   value.ObjectId,
			ResourceId: value.ResourceId,
			Type:       string(value.Type),
		})
	}

	return identities
}

func flattenClusterComputeProfile(input hdinsights.ComputeProfile) []ClusterComputeProfile {
	result := make([]ClusterComputeProfile, 0)
	profile := ClusterComputeProfile{
		Node: flattenClusterNodes(input.Nodes),
	}
	result = append(result, profile)

	return result
}

func flattenClusterNodes(input []hdinsights.NodeProfile) []Node {
	nodes := make([]Node, 0)

	for _, value := range input {
		nodes = append(nodes, Node{
			Count:  int(value.Count),
			VmSize: value.VMSize,
			Type:   value.Type,
		})
	}

	return nodes
}
