package config

import (
	"strings"

	"github.com/zizouhuweidi/retro-station/internal/pkg/config"
	"github.com/zizouhuweidi/retro-station/internal/pkg/config/environemnt"
	typeMapper "github.com/zizouhuweidi/retro-station/internal/pkg/reflection/type_mappper"

	"github.com/iancoleman/strcase"
)

type AppOptions struct {
	DeliveryType string `mapstructure:"deliveryType" env:"DeliveryType"`
	ServiceName  string `mapstructure:"serviceName"  env:"ServiceName"`
}

func NewAppOptions(environment environemnt.Environment) (*AppOptions, error) {
	optionName := strcase.ToLowerCamel(typeMapper.GetTypeNameByT[AppOptions]())
	cfg, err := config.BindConfigKey[*AppOptions](optionName, environment)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func (cfg *AppOptions) GetMicroserviceNameUpper() string {
	return strings.ToUpper(cfg.ServiceName)
}

func (cfg *AppOptions) GetMicroserviceName() string {
	return cfg.ServiceName
}
