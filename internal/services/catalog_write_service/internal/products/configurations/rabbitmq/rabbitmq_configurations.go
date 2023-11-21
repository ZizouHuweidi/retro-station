package rabbitmq

import (
	"github.com/zizouhuweidi/retro-station/internal/pkg/rabbitmq/configurations"
	producerConfigurations "github.com/zizouhuweidi/retro-station/internal/pkg/rabbitmq/producer/configurations"
	createProductIntegrationEvents "github.com/zizouhuweidi/retro-station/internal/services/catalogwriteservice/internal/products/features/creating_product/v1/events/integration_events"
)

func ConfigProductsRabbitMQ(builder configurations.RabbitMQConfigurationBuilder) {
	builder.AddProducer(
		createProductIntegrationEvents.ProductCreatedV1{},
		func(builder producerConfigurations.RabbitMQProducerConfigurationBuilder) {
		})
}
