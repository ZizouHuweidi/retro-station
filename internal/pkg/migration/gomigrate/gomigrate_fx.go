package gomigrate

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/migration"

	"go.uber.org/fx"
)

var (
	// Module provided to fxlog
	// https://uber-go.github.io/fx/modules.html
	Module = fx.Module( //nolint:gochecknoglobals
		"gomigratefx",
		mongoProviders,
	)

	mongoProviders = fx.Provide( //nolint:gochecknoglobals
		migration.ProvideConfig,
		NewGoMigratorPostgres,
	)
)