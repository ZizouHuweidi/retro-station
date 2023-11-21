package config

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/config"
	"github.com/zizouhuweidi/retro-station/internal/pkg/config/environemnt"
	typeMapper "github.com/zizouhuweidi/retro-station/internal/pkg/reflection/type_mappper"

	"github.com/iancoleman/strcase"
)

var optionName = strcase.ToLowerCamel(typeMapper.GetTypeNameByT[GrpcOptions]())

type GrpcOptions struct {
	Port        string `mapstructure:"port"        env:"TcpPort"`
	Host        string `mapstructure:"host"        env:"Host"`
	Development bool   `mapstructure:"development" env:"Development"`
	Name        string `mapstructure:"name"        env:"Name"`
}

func ProvideConfig(environment environemnt.Environment) (*GrpcOptions, error) {
	return config.BindConfigKey[*GrpcOptions](optionName, environment)
}
