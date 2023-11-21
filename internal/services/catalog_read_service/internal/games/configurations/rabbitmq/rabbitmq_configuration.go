package rabbitmq

import (
	"github.com/go-playground/validator"
	"github.com/zizouhuweidi/retro-station/internal/pkg/logger"
	"github.com/zizouhuweidi/retro-station/internal/pkg/messaging/consumer"
	"github.com/zizouhuweidi/retro-station/internal/pkg/otel/tracing"
	rabbitmqConfigurations "github.com/zizouhuweidi/retro-station/internal/pkg/rabbitmq/configurations"
	"github.com/zizouhuweidi/retro-station/internal/pkg/rabbitmq/consumer/configurations"

	createGameExternalEventV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/features/creating_game/v1/events/integration_events/external_events"
	deleteGameExternalEventV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/features/deleting_games/v1/events/integration_events/external_events"
	updateGameExternalEventsV1 "github.com/zizouhuweidi/retro-station/internal/services/catalogreadservice/internal/games/features/updating_games/v1/events/integration_events/external_events"
)

func ConfigGamesRabbitMQ(
	builder rabbitmqConfigurations.RabbitMQConfigurationBuilder,
	logger logger.Logger,
	validator *validator.Validate,
	tracer tracing.AppTracer,
) {
	// add custom message type mappings
	// utils.RegisterCustomMessageTypesToRegistrty(map[string]types.IMessage{"gameCreatedV1": &creatingGameIntegration.GameCreatedV1{}})

	builder.
		AddConsumer(
			createGameExternalEventV1.GameCreatedV1{},
			func(builder configurations.RabbitMQConsumerConfigurationBuilder) {
				builder.WithHandlers(
					func(handlersBuilder consumer.ConsumerHandlerConfigurationBuilder) {
						handlersBuilder.AddHandler(
							createGameExternalEventV1.NewGameCreatedConsumer(
								logger,
								validator,
								tracer,
							),
						)
					},
				)
			}).
		AddConsumer(
			deleteGameExternalEventV1.GameDeletedV1{},
			func(builder configurations.RabbitMQConsumerConfigurationBuilder) {
				builder.WithHandlers(
					func(handlersBuilder consumer.ConsumerHandlerConfigurationBuilder) {
						handlersBuilder.AddHandler(
							deleteGameExternalEventV1.NewGameDeletedConsumer(
								logger,
								validator,
								tracer,
							),
						)
						deleteGameExternalEventV1.NewGameDeletedConsumer(
							logger,
							validator,
							tracer,
						)
					},
				)
			}).
		AddConsumer(
			updateGameExternalEventsV1.GameUpdatedV1{},
			func(builder configurations.RabbitMQConsumerConfigurationBuilder) {
				builder.WithHandlers(
					func(handlersBuilder consumer.ConsumerHandlerConfigurationBuilder) {
						handlersBuilder.AddHandler(
							updateGameExternalEventsV1.NewGameUpdatedConsumer(
								logger,
								validator,
								tracer,
							),
						)
						updateGameExternalEventsV1.NewGameUpdatedConsumer(
							logger,
							validator,
							tracer,
						)
					},
				)
			})
}
