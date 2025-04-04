package monitoring

import (
	"github.com/DataDog/datadog-go/v5/statsd"
	"go.uber.org/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Implementation struct {
	logger   *zap.Logger
	provider config.Provider
	config   MonitoringConfig
	// monitoring providers
	datadogClient *statsd.Client
}

type Params struct {
	fx.In
	Logger    *zap.Logger
	Provider  config.Provider
	Lifecycle fx.Lifecycle
}

type MonitoringConfig struct {
	LoggingEnabled  bool            `yaml:"logging-enabled"`
	LoggingLevel    EmitType        `yaml:"logging-level"`
	MetricsEnabled  bool            `yaml:"metrics-enabled"`
	MetricsLevel    EmitType        `yaml:"metrics-level"`
	MetricsProvider MetricsProvider `yaml:"metrics-provider"`
	DatadogAddr     string          `yaml:"datadog-addr"`
}

type MetricsProvider string

const (
	DATADOG_METRICS_PROVIDER MetricsProvider = "datadog"
)

type EmitType string

const (
	EmitTypeError   EmitType = "error"
	EmitTypeWarn    EmitType = "warn"
	EmitTypeInfo    EmitType = "info"
	EmitTypeSuccess EmitType = "success"
)

type ServiceLayer string

const (
	ModuleServiceLayer     ServiceLayer = "module"
	RepositoryServiceLayer ServiceLayer = "repository"
	ProtocolServiceLayer   ServiceLayer = "protocol"
	EventServiceLayer      ServiceLayer = "event"
)

type EmitRequest struct {
	ActionIdentifier string // used for metrics
	Message          string // used for logs
	EntityIdentifier string
	Layer            ServiceLayer
	LayerSubtype     string
	Type             EmitType
	Data             interface{}
	ExtraData        map[string]string
	Error            error
}
