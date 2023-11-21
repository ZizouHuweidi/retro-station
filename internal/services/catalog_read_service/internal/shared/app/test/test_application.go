package test

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/config/environemnt"
	"github.com/zizouhuweidi/retro-station/internal/pkg/fxapp/test"
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/shared/app"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/shared/configurations/catalogs"

	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

type CatalogsReadTestApplication struct {
	*app.CatalogsReadApplication
	tb fxtest.TB
}

func NewCatalogsReadTestApplication(
	tb fxtest.TB,
	providers []interface{},
	decorates []interface{},
	options []fx.Option,
	logger logger.Logger,
	environment environemnt.Environment,
) *CatalogsReadTestApplication {
	testApp := test.NewTestApplication(
		tb,
		providers,
		decorates,
		options,
		logger,
		environment,
	)

	catalogApplication := &app.CatalogsReadApplication{
		CatalogsServiceConfigurator: catalogs.NewCatalogsServiceConfigurator(testApp),
	}

	return &CatalogsReadTestApplication{
		CatalogsReadApplication: catalogApplication,
		tb:                      tb,
	}
}
