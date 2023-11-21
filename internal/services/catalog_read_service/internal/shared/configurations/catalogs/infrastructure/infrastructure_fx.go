package infrastructure

import (
	"github.com/go-playground/validator"
	"github.com/zizouhuweidi/retro-station/internal/pkg/core"
	"github.com/zizouhuweidi/retro-station/internal/pkg/grpc"
	"github.com/zizouhuweidi/retro-station/internal/pkg/health"
	customEcho "github.com/zizouhuweidi/retro-station/internal/pkg/http/custom_echo"
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger"
	"github.com/zizouhuweidi/retro-station/internal/pkg/mongodb"
	"github.com/zizouhuweidi/retro-station/internal/pkg/otel"
	"github.com/zizouhuweidi/retro-station/internal/pkg/otel/tracing"
	"github.com/zizouhuweidi/retro-station/internal/pkg/rabbitmq"
	"github.com/zizouhuweidi/retro-station/internal/pkg/rabbitmq/configurations"
	"github.com/zizouhuweidi/retro-station/internal/pkg/redis"
	"go.uber.org/fx"

	rabbitmq2 "github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/configurations/rabbitmq"
)

// https://pmihaylov.com/shared-components-go-microservices/
var Module = fx.Module(
	"infrastructurefx",
	// Modules
	core.Module,
	customEcho.Module,
	grpc.Module,
	mongodb.Module,
	otel.Module,
	redis.Module,
	rabbitmq.ModuleFunc(
		func(v *validator.Validate, l logger.Logger, tracer tracing.AppTracer) configurations.RabbitMQConfigurationBuilderFuc {
			return func(builder configurations.RabbitMQConfigurationBuilder) {
				rabbitmq2.ConfigGamesRabbitMQ(builder, l, v, tracer)
			}
		},
	),
	health.Module,

	// Other provides
	fx.Provide(validator.New),
)
