package config

import (
	"time"

	"github.com/metrico/cloki-config/config/reader"
	"github.com/metrico/cloki-config/config/writer"
	"gopkg.in/go-playground/validator.v9"
)

// ============================ BASE ONLY ================================= //
type ClokiBaseDataBase struct {
	User          string `json:"user" mapstructure:"user" default:"cloki_user"`
	Node          string `json:"node" mapstructure:"node" default:"clokinode"`
	Password      string `json:"pass" mapstructure:"pass" default:"cloki_pass"`
	Name          string `json:"name" mapstructure:"name" default:"cloki_data"`
	Host          string `json:"host" mapstructure:"host" default:"127.0.0.1"`
	TableSamples  string `json:"table_samples" mapstructure:"table_samples" default:""`
	TableSeries   string `json:"table_series" mapstructure:"table_series" default:""`
	TableMetrics  string `json:"table_metrics" mapstructure:"table_metrics" default:""`
	Debug         bool   `json:"debug" mapstructure:"debug" default:"false"`
	Port          uint32 `json:"port" mapstructure:"port" default:"9000"`
	ReadTimeout   uint32 `json:"read_timeout" mapstructure:"read_timeout" default:"30"`
	WriteTimeout  uint32 `json:"write_timeout" mapstructure:"write_timeout" default:"30"`
	MaxIdleConn   int    `json:"max_idle_connection" mapstructure:"max_idle_connection" default:"5"`
	MaxOpenConn   int    `json:"max_open_connection" mapstructure:"max_open_connection" default:"50"`
	Primary       bool   `json:"primary" mapstructure:"primary" default:"false"`
	Strategy      string `json:"strategy" mapstructure:"strategy" default:"failover"`
	TTLDays       int    `json:"ttl_days" mapstructure:"ttl_days" default:"7"`
	StoragePolicy string `json:"storage_policy" mapstructure:"storage_policy" default:""`
	Secure        bool   `json:"secure" mapstructure:"secure" default:"false"`
	Cloud         bool   `json:"cloud" mapstructure:"cloud" default:"false"`
	ClusterName   string `json:"cluster_name" mapstructure:"cluster_name" default:""`
	AsyncInsert   bool   `json:"async_insert" mapstructure:"async_insert" default:"false"`
}

type ClokiBaseSettingServer struct {
	ClokiWriter writer.ClokiWriterSettingServer `json:"writer" mapstructure:"writer"`
	ClokiReader reader.ClokiReaderSettingServer `json:"reader" mapstructure:"reader"`

	//Base
	SrartTime                time.Time `default:""`
	FingerPrintType          uint      `default:"1"`
	DataBaseStrategy         uint      `default:"0"`
	CurrentDataBaseIndex     uint      `default:"0"`
	DataDatabaseGroupNodeMap map[string][]string
	Validate                 *validator.Validate
	EnvPrefix                string `default:"QRYN"`
	PluginsPath              string `default:""`

	DATABASE_DATA []ClokiBaseDataBase `json:"database_data" mapstructure:"database_data"`

	SYSTEM_SETTINGS struct {
		HostName             string  `json:"hostname" mapstructure:"hostname" default:"hostname"`
		EnableUserAuditLogin bool    `json:"user_audit_login" mapstructure:"user_audit_login" default:"true"`
		UUID                 string  `json:"uuid" mapstructure:"uuid" default:""`
		DBBulk               int     `json:"db_bulk" mapstructure:"db_bulk" default:"40000"`
		DBTimer              float64 `json:"db_timer" mapstructure:"db_timer" default:"0.2"`
		BufferSizeSample     uint32  `json:"buffer_size_sample" mapstructure:"buffer_size_sample" default:"200000"`
		BufferSizeTimeSeries uint32  `json:"buffer_size_timeseries" mapstructure:"buffer_size_timeseries" default:"200000"`
		ChannelsSample       int     `json:"channels_sample" mapstructure:"channels_sample" default:"2"`
		ChannelsTimeSeries   int     `json:"channels_timeseries" mapstructure:"channels_timeseries" default:"2"`
		HashType             string  `json:"hash_type" mapstructure:"hash_type" default:""`
		CPUMaxProcs          int     `json:"cpu_max_procs" mapstructure:"cpu_max_procs" default:"0"`
	} `json:"system_settings" mapstructure:"system_settings"`

	MULTITENANCE_SETTINGS struct {
		Enabled bool `json:"enabled" mapstructure:"enabled" default:"false"`
	} `json:"multitenance_settings" mapstructure:"multitenance_settings"`

	AUTH_SETTINGS struct {
		AuthTokenHeader string `json:"token_header" mapstructure:"token_header" default:"Auth-Token"`
		AuthTokenExpire uint32 `json:"token_expire" mapstructure:"token_expire" default:"1200"`
	} `json:"auth_settings" mapstructure:"auth_settings"`

	API_SETTINGS struct {
		EnableForceSync   bool `json:"sync_map_force" mapstructure:"sync_map_force" default:"false"`
		EnableTokenAccess bool `json:"enable_token_access" mapstructure:"enable_token_access" default:"true"`
	} `json:"api_settings" mapstructure:"api_settings"`

	HTTP_SETTINGS struct {
		Host          string `json:"host" mapstructure:"host" default:"0.0.0.0"`
		Port          int    `json:"port" mapstructure:"port" default:"3200"`
		ApiPrefix     string `json:"api_prefix" mapstructure:"api_prefix" default:""`
		ApiPromPrefix string `json:"api_prom_prefix" mapstructure:"api_prom_prefix" default:""`
		Prefork       bool   `json:"prefork" mapstructure:"prefork" default:"false"`
		Gzip          bool   `json:"gzip" mapstructure:"gzip" default:"true"`
		GzipStatic    bool   `json:"gzip_static" mapstructure:"gzip_static" default:"true"`
		Debug         bool   `json:"debug" mapstructure:"debug" default:"false"`
		WebSocket     struct {
			Enable bool `json:"enable" mapstructure:"enable" default:"false"`
		} `json:"websocket" mapstructure:"websocket"`
		Enable bool `json:"enable" mapstructure:"enable" default:"true"`
	} `json:"http_settings" mapstructure:"http_settings"`

	HTTPS_SETTINGS struct {
		Host                string `json:"host" mapstructure:"host" default:"0.0.0.0"`
		Port                int    `json:"port" mapstructure:"port" default:"3201"`
		Cert                string `json:"cert" mapstructure:"cert" default:""`
		Key                 string `json:"key" mapstructure:"key" default:""`
		HttpRedirect        bool   `json:"http_redirect" mapstructure:"http_redirect" default:"false"`
		Enable              bool   `json:"enable" mapstructure:"enable" default:"false"`
		MinTLSVersionString string `json:"min_tls_version" mapstructure:"min_tls_version" default:"0"`
		MaxTLSVersionString string `json:"max_tls_version" mapstructure:"max_tls_version" default:"0"`
		MinTLSVersion       uint16 `default:"0"`
		MaxTLSVersion       uint16 `default:"0"`
	} `json:"https_settings" mapstructure:"https_settings"`

	LOG_SETTINGS struct {
		Enable        bool   `json:"enable" mapstructure:"enable" default:"true"`
		MaxAgeDays    uint32 `json:"max_age_days" mapstructure:"max_age_days" default:"7"`
		RotationHours uint32 `json:"rotation_hours" mapstructure:"rotation_hours" default:"24"`
		Path          string `json:"path" mapstructure:"path" default:"./"`
		Level         string `json:"level" mapstructure:"level" default:"error"`
		Name          string `json:"name" mapstructure:"name" default:"ClokiBase.log"`
		Stdout        bool   `json:"stdout" mapstructure:"stdout" default:"false"`
		Json          bool   `json:"json" mapstructure:"json" default:"true"`
		SysLogLevel   string `json:"syslog_level" mapstructure:"syslog_level" default:"LOG_INFO"`
		SysLog        bool   `json:"syslog" mapstructure:"syslog" default:"false"`
		SyslogUri     string `json:"syslog_uri" mapstructure:"syslog_uri" default:""`
	} `json:"log_settings" mapstructure:"log_settings"`

	PROMETHEUS_CLIENT struct {
		PushURL      string   `json:"push_url" mapstructure:"push_url" default:""`
		TargetIP     string   `json:"target_ip" mapstructure:"target_ip" default:""`
		PushInterval uint32   `json:"push_interval" mapstructure:"push_interval" default:"60"`
		PushName     string   `json:"push_name" mapstructure:"push_name" default:""`
		ServiceName  string   `json:"service_name" mapstructure:"service_name" default:"prometheus"`
		MetricsPath  string   `json:"metrics_path" mapstructure:"metrics_path" default:"/metrics"`
		Enable       bool     `json:"enable" mapstructure:"enable" default:"false"`
		AllowIP      []string `json:"allow_ip" mapstructure:"allow_ip" default:"[127.0.0.1]"`
	} `json:"prometheus_client" mapstructure:"prometheus_client"`

	LICENSE_SETTINGS struct {
		RemoteCheckInterval string `json:"remote_check_interval" mapstructure:"remote_check_interval" default:"90d"`
	} `json:"license_settings" mapstructure:"license_settings"`
}
