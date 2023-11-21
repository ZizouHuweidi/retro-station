package mongo

import (
	"context"
	"testing"

	"github.com/zizouhuweidi/retro-station/internal/pkg/logger"
	"github.com/zizouhuweidi/retro-station/internal/pkg/mongodb"
)

var MongoContainerOptionsDecorator = func(t *testing.T, ctx context.Context) interface{} {
	return func(c *mongodb.MongoDbOptions, logger logger.Logger) (*mongodb.MongoDbOptions, error) {
		return NewMongoTestContainers(logger).CreatingContainerOptions(ctx, t)
	}
}
