package options

import "github.com/zizouhuweidi/retro-station/internal/pkg/rabbitmq/types"

type RabbitMQExchangeOptions struct {
	Name       string
	Type       types.ExchangeType
	AutoDelete bool
	Durable    bool
	Args       map[string]any
}
