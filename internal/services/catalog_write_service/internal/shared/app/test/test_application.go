package test

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/config/environemnt"
	"github.com/zizouhuweidi/retro-station/internal/pkg/fxapp/test"
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/shared/app"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/shared/configurations/catalogs"

	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

type CatalogsWriteTestApplication struct {
	*app.CatalogsWriteApplication
	tb fxtest.TB
}

func NewCatalogsWriteTestApplication(
	tb fxtest.TB,
	providers []interface{},
	decorates []interface{},
	options []fx.Option,
	logger logger.Logger,
	environment environemnt.Environment,
) *CatalogsWriteTestApplication {
	testApp := test.NewTestApplication(
		tb,
		providers,
		decorates,
		options,
		logger,
		environment,
	)

	catalogApplication := &app.CatalogsWriteApplication{
		CatalogsServiceConfigurator: catalogs.NewCatalogsServiceConfigurator(testApp),
	}

	return &CatalogsWriteTestApplication{
		CatalogsWriteApplication: catalogApplication,
		tb:                       tb,
	}
}
