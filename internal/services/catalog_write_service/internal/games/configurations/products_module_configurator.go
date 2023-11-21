package configurations

import (
	contracts2 "github.com/zizouhuweidi/retro-station/internal/pkg/fxapp/contracts"
	grpcServer "github.com/zizouhuweidi/retro-station/internal/pkg/grpc"
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger"
	"github.com/zizouhuweidi/retro-station/internal/pkg/messaging/producer"
	"github.com/zizouhuweidi/retro-station/internal/pkg/otel/tracing"
	googleGrpc "google.golang.org/grpc"

	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/configurations/mappings"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/configurations/mediatr"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/contracts/data"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/contracts/params"
	"github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/shared/grpc"
	gamesservice "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/shared/grpc/genproto"
)

type GamesModuleConfigurator struct {
	contracts2.Application
}

func NewGamesModuleConfigurator(
	fxapp contracts2.Application,
) *GamesModuleConfigurator {
	return &GamesModuleConfigurator{
		Application: fxapp,
	}
}

func (c *GamesModuleConfigurator) ConfigureGamesModule() {
	c.ResolveFunc(
		func(logger logger.Logger, uow data.CatalogUnitOfWork, gameRepository data.GameRepository, producer producer.Producer, tracer tracing.AppTracer) error {
			// Config Games Mediators
			err := mediatr.ConfigGamesMediator(logger, uow, gameRepository, producer, tracer)
			if err != nil {
				return err
			}

			// cfg Games Mappings
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
	c.ResolveFunc(func(endpointParams params.GamesEndpointsParams) {
		for _, endpoint := range endpointParams.Endpoints {
			endpoint.MapEndpoint()
		}
	})

	// Config Games Grpc Endpoints
	c.ResolveFunc(
		func(catalogsGrpcServer grpcServer.GrpcServer, grpcService *grpc.GameGrpcServiceServer) error {
			catalogsGrpcServer.GrpcServiceBuilder().RegisterRoutes(func(server *googleGrpc.Server) {
				gamesservice.RegisterGamesServiceServer(server, grpcService)
			})

			return nil
		},
	)
}
