package rabbitmq

import (
	"context"
	"testing"

	"github.com/zizouhuweidi/retro-station/internal/pkg/logger"
	"github.com/zizouhuweidi/retro-station/internal/pkg/rabbitmq/config"
)

var RabbitmqContainerOptionsDecorator = func(t *testing.T, ctx context.Context) interface{} {
	return func(c *config.RabbitmqOptions, logger logger.Logger) (*config.RabbitmqOptions, error) {
		rabbitmqHostOptions, err := NewRabbitMQTestContainers(logger).CreatingContainerOptions(ctx, t)
		c.RabbitmqHostOptions = rabbitmqHostOptions

		return c, err
	}
}
