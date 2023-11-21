package configurations

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/fxapp/contracts"
	logger2 "github.com/zizouhuweidi/retro-station/internal/pkg/logger"
	"github.com/zizouhuweidi/retro-station/internal/pkg/otel/tracing"
	"github.com/zizouhuweidi/retro-station/internal/pkg/web/route"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/products/configurations/mappings"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/products/configurations/mediator"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/products/contracts/data"
)

type ProductsModuleConfigurator struct {
	contracts.Application
}

func NewProductsModuleConfigurator(
	app contracts.Application,
) *ProductsModuleConfigurator {
	return &ProductsModuleConfigurator{
		Application: app,
	}
}

func (c *ProductsModuleConfigurator) ConfigureProductsModule() {
	c.ResolveFunc(
		func(logger logger2.Logger, mongoRepository data.ProductRepository, cacheRepository data.ProductCacheRepository, tracer tracing.AppTracer) error {
			// Config Products Mediators
			err := mediator.ConfigProductsMediator(logger, mongoRepository, cacheRepository, tracer)
			if err != nil {
				return err
			}

			// Config Products Mappings
			err = mappings.ConfigureProductsMappings()
			if err != nil {
				return err
			}
			return nil
		},
	)
}

func (c *ProductsModuleConfigurator) MapProductsEndpoints() {
	// Config Products Http Endpoints
	c.ResolveFuncWithParamTag(func(endpoints []route.Endpoint) {
		for _, endpoint := range endpoints {
			endpoint.MapEndpoint()
		}
	}, `group:"product-routes"`,
	)
}
