package app

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/config/environemnt"
	"github.com/zizouhuweidi/retro-station/internal/pkg/fxapp"
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/shared/configurations/catalogs"

	"go.uber.org/fx"
)

type CatalogsReadApplication struct {
	*catalogs.CatalogsServiceConfigurator
}

func NewCatalogsReadApplication(
	providers []interface{},
	decorates []interface{},
	options []fx.Option,
	logger logger.Logger,
	environment environemnt.Environment,
) *CatalogsReadApplication {
	app := fxapp.NewApplication(providers, decorates, options, logger, environment)
	return &CatalogsReadApplication{
		CatalogsServiceConfigurator: catalogs.NewCatalogsServiceConfigurator(app),
	}
}
