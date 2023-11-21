package configurations

import (
	"reflect"

	types2 "github.com/zizouhuweidi/retro-station/internal/pkg/messaging/types"
	"github.com/zizouhuweidi/retro-station/internal/pkg/messaging/utils"
	"github.com/zizouhuweidi/retro-station/internal/pkg/rabbitmq/producer/options"
	"github.com/zizouhuweidi/retro-station/internal/pkg/rabbitmq/types"
)

type RabbitMQProducerConfiguration struct {
	ProducerMessageType reflect.Type
	ExchangeOptions     *options.RabbitMQExchangeOptions
	RoutingKey          string
	DeliveryMode        uint8
	Priority            uint8
	AppId               string
	Expiration          string
	ReplyTo             string
	ContentEncoding     string
}

func NewDefaultRabbitMQProducerConfiguration(
	messageType types2.IMessage,
) *RabbitMQProducerConfiguration {
	return &RabbitMQProducerConfiguration{
		ExchangeOptions: &options.RabbitMQExchangeOptions{
			Durable: true,
			Type:    types.ExchangeTopic,
			Name:    utils.GetTopicOrExchangeName(messageType),
		},
		DeliveryMode:        2,
		RoutingKey:          utils.GetRoutingKey(messageType),
		ProducerMessageType: utils.GetMessageBaseReflectType(messageType),
	}
}
