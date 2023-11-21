package contracts

import (
	"context"
	"fmt"
	"testing"

	"github.com/zizouhuweidi/retro-station/internal/pkg/core/serializer"
	"github.com/zizouhuweidi/retro-station/internal/pkg/messaging/bus"
	"github.com/zizouhuweidi/retro-station/internal/pkg/rabbitmq/config"
	"github.com/zizouhuweidi/retro-station/internal/pkg/rabbitmq/configurations"
)

type RabbitMQContainerOptions struct {
	Host        string
	VirtualHost string
	Ports       []string
	HostPort    int
	HttpPort    int
	UserName    string
	Password    string
	ImageName   string
	Name        string
	Tag         string
}

func (h *RabbitMQContainerOptions) AmqpEndPoint() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d", h.UserName, h.Password, h.Host, h.HostPort)
}

func (h *RabbitMQContainerOptions) HttpEndPoint() string {
	return fmt.Sprintf("http://%s:%d", h.Host, h.HttpPort)
}

type RabbitMQContainer interface {
	Start(ctx context.Context,
		t *testing.T,
		serializer serializer.EventSerializer,
		rabbitmqBuilderFunc configurations.RabbitMQConfigurationBuilderFuc,
		options ...*RabbitMQContainerOptions) (bus.Bus, error)

	CreatingContainerOptions(
		ctx context.Context,
		t *testing.T,
		options ...*RabbitMQContainerOptions,
	) (*config.RabbitmqHostOptions, error)

	Cleanup(ctx context.Context) error
}
