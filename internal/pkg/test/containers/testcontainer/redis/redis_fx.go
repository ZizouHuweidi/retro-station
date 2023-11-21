package redis

import (
	"context"
	"testing"

	"github.com/zizouhuweidi/retro-station/internal/pkg/logger"
	"github.com/zizouhuweidi/retro-station/internal/pkg/redis"
)

var RedisContainerOptionsDecorator = func(t *testing.T, ctx context.Context) interface{} {
	return func(c *redis.RedisOptions, logger logger.Logger) (*redis.RedisOptions, error) {
		return NewRedisTestContainers(logger).CreatingContainerOptions(ctx, t)
	}
}
