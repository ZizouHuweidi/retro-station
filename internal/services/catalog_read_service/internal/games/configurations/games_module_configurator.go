package configurations

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/fxapp/contracts"
	logger2 "github.com/zizouhuweidi/retro-station/internal/pkg/logger"
	"github.com/zizouhuweidi/retro-station/internal/pkg/otel/tracing"
	"github.com/zizouhuweidi/retro-station/internal/pkg/web/route"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/configurations/mappings"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/configurations/mediator"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/contracts/data"
)

type GamesModuleConfigurator struct {
	contracts.Application
}

func NewGamesModuleConfigurator(
	app contracts.Application,
) *GamesModuleConfigurator {
	return &GamesModuleConfigurator{
		Application: app,
	}
}

func (c *GamesModuleConfigurator) ConfigureGamesModule() {
	c.ResolveFunc(
		func(logger logger2.Logger, mongoRepository data.GameRepository, cacheRepository data.GameCacheRepository, tracer tracing.AppTracer) error {
			// Config Games Mediators
			err := mediator.ConfigGamesMediator(logger, mongoRepository, cacheRepository, tracer)
			if err != nil {
				return err
			}

			// Config Games Mappings
			err = mappings.ConfigureGamesMappings()
			if err != nil {
				return err
			}
			return nil
		},
	)
}

func (c *GamesModuleConfigurator) MapGamesEndpoints() {
	// Config Games Http Endpoints
	c.ResolveFuncWithParamTag(func(endpoints []route.Endpoint) {
		for _, endpoint := range endpoints {
			endpoint.MapEndpoint()
		}
	}, `group:"game-routes"`,
	)
}
