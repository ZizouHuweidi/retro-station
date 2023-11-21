package configurations

import (
	consumerConfigurations "github.com/zizouhuweidi/retro-station/internal/pkg/rabbitmq/consumer/configurations"
	producerConfigurations "github.com/zizouhuweidi/retro-station/internal/pkg/rabbitmq/producer/configurations"
)

type RabbitMQConfiguration struct {
	ProducersConfigurations []*producerConfigurations.RabbitMQProducerConfiguration
	ConsumersConfigurations []*consumerConfigurations.RabbitMQConsumerConfiguration
}
