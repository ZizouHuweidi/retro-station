package configurations

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/es/contracts/store"
	contracts2 "github.com/zizouhuweidi/retro-station/internal/pkg/fxapp/contracts"
	grpcServer "github.com/zizouhuweidi/retro-station/internal/pkg/grpc"
	customEcho "github.com/zizouhuweidi/retro-station/internal/pkg/http/custom_echo"
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger"
	"github.com/zizouhuweidi/retro-station/internal/pkg/otel/tracing"
	"github.com/zizouhuweidi/retro-station/internal/pkg/web/route"
	"github.com/zizouhuweidi/retro-station/internal/services/orderservice/internal/orders/configurations/mappings"
	"github.com/zizouhuweidi/retro-station/internal/services/orderservice/internal/orders/configurations/mediatr"
	"github.com/zizouhuweidi/retro-station/internal/services/orderservice/internal/orders/contracts/repositories"
	"github.com/zizouhuweidi/retro-station/internal/services/orderservice/internal/orders/models/orders/aggregate"
	"github.com/zizouhuweidi/retro-station/internal/services/orderservice/internal/shared/contracts"
	"github.com/zizouhuweidi/retro-station/internal/services/orderservice/internal/shared/grpc"
	ordersservice "github.com/zizouhuweidi/retro-station/internal/services/orderservice/internal/shared/grpc/genproto"

	"github.com/go-playground/validator"
	googleGrpc "google.golang.org/grpc"
)

type OrdersModuleConfigurator struct {
	contracts2.Application
}

func NewOrdersModuleConfigurator(
	app contracts2.Application,
) *OrdersModuleConfigurator {
	return &OrdersModuleConfigurator{
		Application: app,
	}
}

func (c *OrdersModuleConfigurator) ConfigureOrdersModule() {
	c.ResolveFunc(
		func(logger logger.Logger,
			server customEcho.EchoHttpServer,
			orderRepository repositories.OrderMongoRepository,
			orderAggregateStore store.AggregateStore[*aggregate.Order],
			tracer tracing.AppTracer,
		) error {
			// Config Orders Mappings
			err := mappings.ConfigureOrdersMappings()
			if err != nil {
				return err
			}

			// Config Orders Mediators
			err = mediatr.ConfigOrdersMediator(logger, orderRepository, orderAggregateStore, tracer)
			if err != nil {
				return err
			}

			return nil
		},
	)
}

func (c *OrdersModuleConfigurator) MapOrdersEndpoints() {
	// Config Orders Http Endpoints
	c.ResolveFuncWithParamTag(func(endpoints []route.Endpoint) {
		for _, endpoint := range endpoints {
			endpoint.MapEndpoint()
		}
	}, `group:"order-routes"`,
	)

	// Config Orders Grpc Endpoints
	c.ResolveFunc(
		func(ordersGrpcServer grpcServer.GrpcServer, ordersMetrics *contracts.OrdersMetrics, logger logger.Logger, validator *validator.Validate) error {
			orderGrpcService := grpc.NewOrderGrpcService(logger, validator, ordersMetrics)
			ordersGrpcServer.GrpcServiceBuilder().RegisterRoutes(func(server *googleGrpc.Server) {
				ordersservice.RegisterOrdersServiceServer(server, orderGrpcService)
			})
			return nil
		},
	)
}
