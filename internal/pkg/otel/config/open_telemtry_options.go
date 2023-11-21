package config

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/config"
	"github.com/zizouhuweidi/retro-station/internal/pkg/config/environemnt"
	typeMapper "github.com/zizouhuweidi/retro-station/internal/pkg/reflection/type_mappper"

	"github.com/iancoleman/strcase"
)

var optionName = strcase.ToLowerCamel(typeMapper.GetTypeNameByT[OpenTelemetryOptions]())

type OpenTelemetryOptions struct {
	Enabled               bool                   `mapstructure:"enabled"`
	ServiceName           string                 `mapstructure:"serviceName"`
	InstrumentationName   string                 `mapstructure:"instrumentationName"`
	Id                    int64                  `mapstructure:"id"`
	AlwaysOnSampler       bool                   `mapstructure:"alwaysOnSampler"`
	JaegerExporterOptions *JaegerExporterOptions `mapstructure:"jaegerExporterOptions"`
	ZipkinExporterOptions *ZipkinExporterOptions `mapstructure:"zipkinExporterOptions"`
	OTelMetricsOptions    *OTelMetricsOptions    `mapstructure:"otelMetricsOptions"`
	UseStdout             bool                   `mapstructure:"useStdout"`
}

type JaegerExporterOptions struct {
	AgentHost string `mapstructure:"agentHost"`
	AgentPort string `mapstructure:"agentPort"`
}

type ZipkinExporterOptions struct {
	Url string `mapstructure:"url"`
}

type OTelMetricsOptions struct {
	Host             string `mapstructure:"host"`
	Port             string `mapstructure:"port"`
	Name             string `mapstructure:"name"`
	MetricsRoutePath string `mapstructure:"metricsRoutePath"`
}

func ProvideOtelConfig(environment environemnt.Environment) (*OpenTelemetryOptions, error) {
	return config.BindConfigKey[*OpenTelemetryOptions](optionName, environment)
}
