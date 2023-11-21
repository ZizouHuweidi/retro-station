package infrastructure

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/core"
	gormPostgres "github.com/zizouhuweidi/retro-station/internal/pkg/gorm_postgres"
	"github.com/zizouhuweidi/retro-station/internal/pkg/grpc"
	"github.com/zizouhuweidi/retro-station/internal/pkg/health"
	customEcho "github.com/zizouhuweidi/retro-station/internal/pkg/http/custom_echo"
	"github.com/zizouhuweidi/retro-station/internal/pkg/otel"
	"github.com/zizouhuweidi/retro-station/internal/pkg/rabbitmq"
	"github.com/zizouhuweidi/retro-station/internal/pkg/rabbitmq/configurations"
	rabbitmq2 "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/products/configurations/rabbitmq"

	"github.com/go-playground/validator"
	"go.uber.org/fx"
)

// https://pmihaylov.com/shared-components-go-microservices/

var Module = fx.Module(
	"infrastructurefx",
	// Modules
	core.Module,
	customEcho.Module,
	grpc.Module,
	gormPostgres.Module,
	otel.Module,
	rabbitmq.ModuleFunc(
		func() configurations.RabbitMQConfigurationBuilderFuc {
			return func(builder configurations.RabbitMQConfigurationBuilder) {
				rabbitmq2.ConfigProductsRabbitMQ(builder)
			}
		},
	),
	health.Module,

	// Other provides
	fx.Provide(validator.New),
)