package migration

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/config"
	"github.com/zizouhuweidi/retro-station/internal/pkg/config/environemnt"
	typeMapper "github.com/zizouhuweidi/retro-station/internal/pkg/reflection/type_mappper"

	"github.com/iancoleman/strcase"
)

type CommandType string

const (
	Up   CommandType = "up"
	Down CommandType = "down"
)

type MigrationOptions struct {
	Host          string `mapstructure:"host"`
	Port          int    `mapstructure:"port"`
	User          string `mapstructure:"user"`
	DBName        string `mapstructure:"dbName"`
	SSLMode       bool   `mapstructure:"sslMode"`
	Password      string `mapstructure:"password"`
	VersionTable  string `mapstructure:"versionTable"`
	MigrationsDir string `mapstructure:"migrationsDir"`
	SkipMigration bool   `mapstructure:"skipMigration"`
}

func ProvideConfig(environment environemnt.Environment) (*MigrationOptions, error) {
	optionName := strcase.ToLowerCamel(typeMapper.GetTypeNameByT[MigrationOptions]())

	return config.BindConfigKey[*MigrationOptions](optionName, environment)
}
