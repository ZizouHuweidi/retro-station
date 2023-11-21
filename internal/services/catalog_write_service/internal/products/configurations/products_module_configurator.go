package configurations

import (
	contracts2 "github.com/zizouhuweidi/retro-station/internal/pkg/fxapp/contracts"
	grpcServer "github.com/zizouhuweidi/retro-station/internal/pkg/grpc"
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger"
	"github.com/zizouhuweidi/retro-station/internal/pkg/messaging/producer"
	"github.com/zizouhuweidi/retro-station/internal/pkg/otel/tracing"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/products/configurations/mappings"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/products/configurations/mediatr"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/products/contracts/data"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/products/contracts/params"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/shared/grpc"
	productsservice "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/shared/grpc/genproto"

	googleGrpc "google.golang.org/grpc"
)

type ProductsModuleConfigurator struct {
	contracts2.Application
}

func NewProductsModuleConfigurator(
	fxapp contracts2.Application,
) *ProductsModuleConfigurator {
	return &ProductsModuleConfigurator{
		Application: fxapp,
	}
}

func (c *ProductsModuleConfigurator) ConfigureProductsModule() {
	c.ResolveFunc(
		func(logger logger.Logger, uow data.CatalogUnitOfWork, productRepository data.ProductRepository, producer producer.Producer, tracer tracing.AppTracer) error {
			// Config Products Mediators
			err := mediatr.ConfigProductsMediator(logger, uow, productRepository, producer, tracer)
			if err != nil {
				return err
			}

			// cfg Products Mappings
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
	c.ResolveFunc(func(endpointParams params.ProductsEndpointsParams) {
		for _, endpoint := range endpointParams.Endpoints {
			endpoint.MapEndpoint()
		}
	})

	// Config Products Grpc Endpoints
	c.ResolveFunc(
		func(catalogsGrpcServer grpcServer.GrpcServer, grpcService *grpc.ProductGrpcServiceServer) error {
			catalogsGrpcServer.GrpcServiceBuilder().RegisterRoutes(func(server *googleGrpc.Server) {
				productsservice.RegisterProductsServiceServer(server, grpcService)
			})

			return nil
		},
	)
}
