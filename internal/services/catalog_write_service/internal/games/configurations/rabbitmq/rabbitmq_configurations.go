package rabbitmq

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/rabbitmq/configurations"
	producerConfigurations "github.com/zizouhuweidi/retro-station/internal/pkg/rabbitmq/producer/configurations"

	createGameIntegrationEvents "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/games/features/creating_game/v1/events/integration_events"
)

func ConfigGamesRabbitMQ(builder configurations.RabbitMQConfigurationBuilder) {
	builder.AddProducer(
		createGameIntegrationEvents.GameCreatedV1{},
		func(builder producerConfigurations.RabbitMQProducerConfigurationBuilder) {
		})
}
