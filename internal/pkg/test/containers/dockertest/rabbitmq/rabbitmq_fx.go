package rabbitmq

import (
	"context"
	"testing"

	"github.com/zizouhuweidi/retro-station/internal/pkg/logger"
	"github.com/zizouhuweidi/retro-station/internal/pkg/rabbitmq/config"
)

var RabbitmqDockerTestContainerOptionsDecorator = func(t *testing.T, ctx context.Context) interface{} {
	return func(c *config.RabbitmqOptions, logger logger.Logger) (*config.RabbitmqOptions, error) {
		rabbitmqHostOptions, err := NewRabbitMQDockerTest(logger).CreatingContainerOptions(ctx, t)
		c.RabbitmqHostOptions = rabbitmqHostOptions

		return c, err
	}
}
